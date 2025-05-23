package handler

import (
	"bytes"
	"context"
	"path/filepath"
	"strconv"

	"github.com/aiagt/aiagt/tools/utils/logger"

	"github.com/bytedance/gopkg/util/gopool"

	"github.com/aiagt/aiagt/apps/knowledge/model"
	"github.com/aiagt/aiagt/common/cos"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/cloudwego/eino-ext/components/document/parser/html"
	"github.com/cloudwego/eino-ext/components/document/parser/pdf"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/recursive"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"code.sajari.com/docconv"
)

var (
	documentProcessPoolSize int32 = 128
	documentProcessPool           = gopool.NewPool("document_process", documentProcessPoolSize, gopool.NewConfig())
)

// UploadDocument implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) UploadDocument(ctx context.Context, req *knowledgesvc.UploadDocumentReq) (resp *base.Empty, err error) {
	knowledgeID := req.GetKnowledgeId()

	knowledge, err := s.knowledgeDao.GetByID(ctx, knowledgeID)
	if err != nil {
		return nil, bizUploadDocument.NewErr(err).Log(ctx, "get knowledge by id error")
	}

	// upload cos
	filename := uuid.New().String() + req.GetType()
	path := filepath.Join(cos.KnowledgeDir, filename)

	_, err = cos.Cli().Object.Put(ctx, path, bytes.NewReader(req.GetContent()), cos.WithContentTypeByExt(req.GetType()))
	if err != nil {
		return nil, bizUploadDocument.NewErr(err).Log(ctx, "upload document error")
	}

	// insert document to mysql
	knowledgeDocument := model.KnowledgeDocument{
		KnowledgeID: req.GetKnowledgeId(),
		Name:        req.GetName(),
		Url:         path,
		Type:        req.GetType(),
		Preview:     utils.TruncateString(string(req.GetContent()), 200),
		Size:        int64(len(req.GetContent())),
		ChunkSize:   req.GetChunkSize(),
		OverlapSize: req.GetOverlapSize(),
		Separators:  req.GetSeparators(),
		Status:      knowledgesvc.EmbeddingStatus_PENDING,
	}

	err = s.documentDao.Create(ctx, &knowledgeDocument)
	if err != nil {
		return nil, bizUploadDocument.NewErr(err).Log(ctx, "create knowledge document error")
	}

	// parse and split document
	parsedDocuments, err := s.parseDocument(ctx, req.GetType(), req.GetContent())
	if err != nil {
		return nil, bizUploadDocument.NewErr(err).Log(ctx, "parse document error")
	}

	splitDocuments, err := s.splitDocument(ctx, parsedDocuments, req.GetChunkSize(), req.GetOverlapSize(), req.GetSeparators())
	if err != nil {
		return nil, bizUploadDocument.NewErr(err).Log(ctx, "split document error")
	}

	documentProcessPool.Go(func() {
		var (
			ctx          = context.Background()
			err          error
			processError error
		)

		defer func() {
			if processError != nil {
				err = s.documentDao.UpdateStatus(ctx, knowledgeDocument.ID, knowledgesvc.EmbeddingStatus_FAILED, processError.Error())
			} else {
				err = s.documentDao.UpdateStatus(ctx, knowledgeDocument.ID, knowledgesvc.EmbeddingStatus_SUCCESS, "")
			}

			if err != nil {
				logger.Errorf("process knowledge document error: %v", err)
			}
		}()

		var knowledgeChunks []*model.KnowledgeChunk

		for i, doc := range splitDocuments {
			knowledgeChunks = append(knowledgeChunks, &model.KnowledgeChunk{
				KnowledgeID: knowledge.ID,
				DocumentID:  knowledgeDocument.ID,
				Content:     doc.Content,
				Order:       int64(i + 1),
				Status:      knowledgesvc.EmbeddingStatus_PENDING,
			})

			if doc.MetaData == nil {
				doc.MetaData = make(map[string]interface{})
			}
		}

		// insert document chunk to mysql
		err = s.chunkDao.CreateBatch(ctx, knowledgeChunks)
		if err != nil {
			processError = err
			return
		}

		// update chunk metadata
		for i, doc := range splitDocuments {
			knowledgeChunk := knowledgeChunks[i]

			doc.ID = strconv.FormatInt(knowledgeChunk.ID, 10)
			doc.MetaData["knowledge_id"] = strconv.FormatInt(knowledge.ID, 10)
			doc.MetaData["document_id"] = strconv.FormatInt(knowledgeDocument.ID, 10)
		}

		// embedding and insert to milvus
		indexer, err := s.indexer(ctx)
		if err != nil {
			processError = err
			return
		}

		_, err = indexer.Store(ctx, splitDocuments)
		if err != nil {
			processError = err
			return
		}
	})

	return
}

func (s *KnowledgeServiceImpl) parseDocument(ctx context.Context, docType string, content []byte) ([]*schema.Document, error) {
	documents := make([]*schema.Document, 0)

	switch docType {
	case ".pdf", "application/pdf":
		parser, err := pdf.NewPDFParser(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, "new html parser error")
		}

		documents, err = parser.Parse(ctx, bytes.NewReader(content))
		if err != nil {
			return nil, errors.Wrap(err, "parse html document error")
		}
	case ".docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		res, err := docconv.Convert(bytes.NewReader(content), "application/vnd.openxmlformats-officedocument.wordprocessingml.document", true)
		if err != nil {
			return nil, errors.Wrap(err, "parse word document error")
		}
		documents = append(documents, &schema.Document{Content: res.Body})
	case ".doc", "application/msword":
		res, err := docconv.Convert(bytes.NewReader(content), "application/msword", true)
		if err != nil {
			return nil, errors.Wrap(err, "parse word document error")
		}
		documents = append(documents, &schema.Document{Content: res.Body})
	case ".html", "text/html":
		parser, err := html.NewParser(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, "new html parser error")
		}

		documents, err = parser.Parse(ctx, bytes.NewReader(content))
		if err != nil {
			return nil, errors.Wrap(err, "parse html document error")
		}
	default:
		documents = []*schema.Document{{Content: string(content)}}
	}

	return documents, nil
}

func (s *KnowledgeServiceImpl) splitDocument(ctx context.Context, src []*schema.Document, chunkSize, overlapSize int64, separators []string) ([]*schema.Document, error) {
	splitter, err := recursive.NewSplitter(ctx, &recursive.Config{
		ChunkSize:   int(chunkSize),
		OverlapSize: int(overlapSize),
		Separators:  separators,
		KeepType:    recursive.KeepTypeEnd,
	})
	if err != nil {
		return nil, errors.Wrap(err, "new splitter error")
	}

	result, err := splitter.Transform(ctx, src)
	if err != nil {
		return nil, errors.Wrap(err, "split document error")
	}

	return result, nil
}

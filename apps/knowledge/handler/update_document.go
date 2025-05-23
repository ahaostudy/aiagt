package handler

import (
	"context"
	"io"
	"strconv"

	"github.com/aiagt/aiagt/apps/knowledge/model"
	"github.com/aiagt/aiagt/common/cos"
	"github.com/aiagt/aiagt/tools/utils/logger"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
)

// UpdateDocument implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) UpdateDocument(ctx context.Context, req *knowledgesvc.UpdateDocumentReq) (resp *base.Empty, err error) {
	knowledgeDocument, err := s.documentDao.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, bizUpdateDocument.NewErr(err).Log(ctx, "get knowledge document by id error")
	}

	knowledgeDocument.ChunkSize = req.GetChunkSize()
	knowledgeDocument.OverlapSize = req.GetOverlapSize()
	knowledgeDocument.Separators = req.GetSeparators()
	knowledgeDocument.Status = knowledgesvc.EmbeddingStatus_PENDING

	err = s.documentDao.Save(ctx, knowledgeDocument)
	if err != nil {
		return nil, bizUploadDocument.NewErr(err).Log(ctx, "create knowledge document error")
	}

	fileResp, err := cos.Cli().Object.Get(ctx, knowledgeDocument.Url, nil)
	if err != nil {
		return nil, bizUploadDocument.NewErr(err).Log(ctx, "get knowledge document file error")
	}
	defer fileResp.Body.Close()

	fileBody, err := io.ReadAll(fileResp.Body)
	if err != nil {
		return nil, bizUploadDocument.NewErr(err).Log(ctx, "read knowledge document file error")
	}

	// parse and split document
	parsedDocuments, err := s.parseDocument(ctx, knowledgeDocument.Type, fileBody)
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
				KnowledgeID: knowledgeDocument.KnowledgeID,
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
			doc.MetaData["knowledge_id"] = strconv.FormatInt(knowledgeDocument.KnowledgeID, 10)
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

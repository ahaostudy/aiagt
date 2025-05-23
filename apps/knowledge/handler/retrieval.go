package handler

import (
	"context"
	"fmt"
	"strconv"

	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
	"github.com/aiagt/aiagt/pkg/lists"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
	"github.com/cloudwego/eino/schema"
)

// Retrieval implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) Retrieval(ctx context.Context, req *knowledgesvc.RetrievalReq) (resp *knowledgesvc.RetrievalResp, err error) {
	knowledgeList, err := s.knowledgeDao.GetByIDs(ctx, req.GetKnowledgeIds())
	if err != nil {
		return nil, bizRetrieval.NewErr(err).Log(ctx, "get knowledge by id error")
	}

	var retrievalChunks []*schema.Document

	// retrieval document chunks
	for _, knowledge := range knowledgeList {
		retriever, err := s.retriever(ctx, int(knowledge.TopK), knowledge.ScoreThreshold)
		if err != nil {
			return nil, bizRetrieval.NewErr(err).Log(ctx, "new retriever error")
		}

		documents, err := retriever.Retrieve(ctx, req.GetQuery(), milvus.WithFilter(fmt.Sprintf("knowledge_id == '%d'", knowledge.ID)))
		if err != nil {
			return nil, bizRetrieval.NewErr(err).Log(ctx, "retriever query error")
		}

		retrievalChunks = append(retrievalChunks, documents...)
	}

	// build response
	resp = &knowledgesvc.RetrievalResp{
		Chunks: lists.Map(retrievalChunks, func(t *schema.Document) *knowledgesvc.KnowledgeChunk {
			knowledgeIDStr, _ := t.MetaData["knowledge_id"].(string)
			documentIDStr, _ := t.MetaData["document_id"].(string)

			knowledgeID, _ := strconv.ParseInt(knowledgeIDStr, 10, 64)
			documentID, _ := strconv.ParseInt(documentIDStr, 10, 64)

			return &knowledgesvc.KnowledgeChunk{
				Id:          t.ID,
				KnowledgeId: knowledgeID,
				DocumentId:  documentID,
				Content:     t.Content,
			}
		}),
	}

	return
}

package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/baseutil"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	"gorm.io/gorm"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
)

// GetDocumentByKnowledge implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) GetDocumentByKnowledge(ctx context.Context, req *base.IDReq) (resp *knowledgesvc.GetDocumentByKnowledgeResp, err error) {
	knowledge, err := s.knowledgeDao.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, bizGetDocumentByKnowledge.NewErr(err).Log(ctx, "get knowledge by id error")
	}

	if ctxutil.Forbidden(ctx, knowledge.AuthorID) {
		return nil, bizGetDocumentByKnowledge.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
	}

	documents, err := s.documentDao.ListBy(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("knowledge_id = ?", req.GetId())
	})
	if err != nil {
		return nil, bizGetDocumentByKnowledge.NewErr(err).Log(ctx, "get document by knowledge")
	}

	resp = new(knowledgesvc.GetDocumentByKnowledgeResp)

	for _, document := range documents {
		resp.Documents = append(resp.Documents, &knowledgesvc.KnowledgeDocument{
			Id:             document.ID,
			KnowledgeId:    document.KnowledgeID,
			Name:           document.Name,
			Url:            document.Url,
			Type:           document.Type,
			Preview:        document.Preview,
			Size:           document.Size,
			RetrievalCount: document.RetrievalCount,
			UpdatedAt:      baseutil.NewBaseTime(document.UpdatedAt),
			ChunkSize:      document.ChunkSize,
			OverlapSize:    document.OverlapSize,
			Separators:     document.Separators,
			Status:         document.Status,
		})
	}

	return
}

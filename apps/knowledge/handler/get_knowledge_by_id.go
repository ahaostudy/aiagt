package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/baseutil"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/pkg/utils"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
)

// GetKnowledgeByID implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) GetKnowledgeByID(ctx context.Context, req *base.IDReq) (resp *knowledgesvc.Knowledge, err error) {
	knowledge, err := s.knowledgeDao.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, bizGetKnowledgeByID.NewErr(err).Log(ctx, "get knowledge by id error")
	}

	if ctxutil.Forbidden(ctx, knowledge.AuthorID) {
		return nil, bizGetKnowledgeByID.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
	}

	resp = &knowledgesvc.Knowledge{
		Id:             knowledge.ID,
		Name:           knowledge.Name,
		Description:    knowledge.Description,
		AuthorId:       knowledge.AuthorID,
		Author:         nil,
		IsPrivate:      knowledge.IsPrivate,
		Logo:           knowledge.Logo,
		EmbedModelId:   knowledge.EmbedModelID,
		TopK:           knowledge.TopK,
		ScoreThreshold: utils.OPtrOf(knowledge.ScoreThreshold),
		CreatedAt:      baseutil.NewBaseTime(knowledge.CreatedAt),
		UpdatedAt:      baseutil.NewBaseTime(knowledge.UpdatedAt),
		PublishedAt:    baseutil.NewBaseTimeP(knowledge.PublishedAt),
	}

	return
}

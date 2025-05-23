package handler

import (
	"context"
	"time"

	"github.com/aiagt/aiagt/apps/knowledge/model"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
)

// SaveKnowledge implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) SaveKnowledge(ctx context.Context, req *knowledgesvc.SaveKnowledgeReq) (resp *base.Empty, err error) {
	knowledge := new(model.Knowledge)

	if req.GetId() != 0 {
		knowledge, err = s.knowledgeDao.GetByID(ctx, req.GetId())
		if err != nil {
			return nil, bizSaveKnowledge.NewErr(err).Log(ctx, "get knowledge by id error")
		}

		if ctxutil.Forbidden(ctx, knowledge.AuthorID) {
			return nil, bizSaveKnowledge.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
		}
	}

	knowledge.Name = req.GetName()
	knowledge.Description = req.GetDescription()
	knowledge.AuthorID = ctxutil.UserID(ctx)
	knowledge.IsPrivate = req.GetIsPrivate()
	knowledge.EmbedModelID = req.GetEmbedModelId()
	knowledge.TopK = req.GetTopK()
	knowledge.ScoreThreshold = req.GetScoreThreshold()
	knowledge.Logo = req.GetLogo()
	knowledge.UpdatedAt = time.Now()

	err = s.knowledgeDao.Save(ctx, knowledge)
	if err != nil {
		return nil, err
	}

	return
}

package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/knowledge/model"
	"github.com/aiagt/aiagt/common/baseutil"
	"github.com/aiagt/aiagt/common/ctxutil"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
	"github.com/aiagt/aiagt/pkg/lists"
	"github.com/aiagt/aiagt/pkg/utils"
)

// ListKnowledge implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) ListKnowledge(ctx context.Context, req *knowledgesvc.ListKnowledgeReq) (resp *knowledgesvc.ListKnowledgeResp, err error) {
	userID := ctxutil.UserID(ctx)

	knowledgeList, page, err := s.knowledgeDao.List(ctx, req, userID)
	if err != nil {
		return nil, bizListKnowledge.NewErr(err)
	}

	resp = &knowledgesvc.ListKnowledgeResp{
		KnowledgeList: lists.Map(knowledgeList, func(t *model.Knowledge) *knowledgesvc.Knowledge {
			return &knowledgesvc.Knowledge{
				Id:             t.ID,
				Name:           t.Name,
				Description:    t.Description,
				AuthorId:       t.AuthorID,
				IsPrivate:      t.IsPrivate,
				Logo:           t.Logo,
				EmbedModelId:   t.EmbedModelID,
				TopK:           t.TopK,
				ScoreThreshold: utils.OPtrOf(t.ScoreThreshold),
				CreatedAt:      baseutil.NewBaseTime(t.CreatedAt),
				UpdatedAt:      baseutil.NewBaseTime(t.UpdatedAt),
				PublishedAt:    baseutil.NewBaseTimeP(t.PublishedAt),
			}
		}),
		Pagination: page,
	}
	return
}

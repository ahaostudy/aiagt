package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/baseutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	knowledgesvc "github.com/aiagt/aiagt/kitex_gen/knowledgesvc"
	"github.com/aiagt/aiagt/pkg/utils"
)

// GetKnowledgeByIDs implements the KnowledgeServiceImpl interface.
func (s *KnowledgeServiceImpl) GetKnowledgeByIDs(ctx context.Context, req *base.IDsReq) (resp []*knowledgesvc.Knowledge, err error) {
	knowledgeList, err := s.knowledgeDao.GetByIDs(ctx, req.GetIds())
	if err != nil {
		return nil, bizGetKnowledgeByIDs.NewErr(err).Log(ctx, "get knowledge by ids err")
	}

	resp = make([]*knowledgesvc.Knowledge, 0, len(knowledgeList))

	for _, knowledge := range knowledgeList {
		resp = append(resp, &knowledgesvc.Knowledge{
			Id:             knowledge.ID,
			Name:           knowledge.Name,
			Description:    knowledge.Description,
			AuthorId:       knowledge.AuthorID,
			IsPrivate:      knowledge.IsPrivate,
			EmbedModelId:   knowledge.EmbedModelID,
			TopK:           knowledge.TopK,
			ScoreThreshold: utils.OPtrOf(knowledge.ScoreThreshold),
			Logo:           knowledge.Logo,
			CreatedAt:      baseutil.NewBaseTime(knowledge.CreatedAt),
			UpdatedAt:      baseutil.NewBaseTime(knowledge.UpdatedAt),
			PublishedAt:    baseutil.NewBaseTimeP(knowledge.PublishedAt),
		})
	}

	return
}

package handler

import (
	"context"

	"github.com/aiagt/aiagt/apps/workflow/model"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc"
)

// CreateWorkflow implements the WorkflowServiceImpl interface.
func (s *WorkflowServiceImpl) CreateWorkflow(ctx context.Context, req *workflowsvc.CreateWorkflowReq) (resp *base.Empty, err error) {
	userID, ok := ctxutil.GetUserID(ctx)
	if !ok {
		return nil, bizCreateWorkflow.NewErr(err).Log(ctx, "get user id error")
	}

	workflow := model.Workflow{
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		DescriptionMd: req.GetDescriptionMd(),
		AuthorID:      userID,
		IsPrivate:     req.GetIsPrivate(),
		Logo:          req.GetLogo(),
	}

	err = s.workflowDao.Create(ctx, &workflow)
	if err != nil {
		return nil, bizCreateWorkflow.NewErr(err).Log(ctx, "create workflow error")
	}

	return
}

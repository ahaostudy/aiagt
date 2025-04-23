package handler

import (
	"context"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeleteWorkflowNode implements the WorkflowServiceImpl interface.
func (s *WorkflowServiceImpl) DeleteWorkflowNode(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	workflowNode, err := s.workflowNodeDao.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, bizDeleteWorkflowNode.NewErr(err).Log(ctx, "get workflow node error", err)
	}

	workflow, err := s.workflowDao.GetByID(ctx, workflowNode.WorkflowID)
	if err != nil {
		return nil, bizDeleteWorkflowNode.NewErr(err).Log(ctx, "get workflow error", err)
	}

	if ctxutil.Forbidden(ctx, workflow.AuthorID) {
		return nil, bizDeleteWorkflowNode.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
	}

	err = s.workflowNodeDao.Delete(ctx, req.GetId())
	if err != nil {
		return nil, bizDeleteWorkflowNode.NewErr(err).Log(ctx, "delete workflow node error", err)
	}

	return
}

package handler

import (
	"context"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// DeleteWorkflow implements the WorkflowServiceImpl interface.
func (s *WorkflowServiceImpl) DeleteWorkflow(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	workflow, err := s.workflowDao.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, bizDeleteWorkflow.NewErr(err).Log(ctx, "delete workflow error")
	}

	if ctxutil.Forbidden(ctx, workflow.AuthorID) {
		return nil, bizDeleteWorkflow.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
	}

	err = s.workflowDao.Delete(ctx, req.GetId())
	if err != nil {
		return nil, bizDeleteWorkflow.NewErr(err).Log(ctx, "delete workflow error")
	}

	return
}

package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc"
)

// GetWorkflowByID implements the WorkflowServiceImpl interface.
func (s *WorkflowServiceImpl) GetWorkflowByID(ctx context.Context, req *base.IDReq) (resp *workflowsvc.Workflow, err error) {
	return
}

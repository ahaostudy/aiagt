package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc"
)

// UpdateWorkflow implements the WorkflowServiceImpl interface.
func (s *WorkflowServiceImpl) UpdateWorkflow(ctx context.Context, req *workflowsvc.UpdateWorkflowReq) (resp *base.Empty, err error) {
	return
}

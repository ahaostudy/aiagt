package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc"
)

// UpdateWorkflowNode implements the WorkflowServiceImpl interface.
func (s *WorkflowServiceImpl) UpdateWorkflowNode(ctx context.Context, req *workflowsvc.UpdateWorkflowNodeReq) (resp *base.Empty, err error) {
	return
}

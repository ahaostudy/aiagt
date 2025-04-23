package handler

import (
	"context"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc"
)

// GetWorkflowNodeByID implements the WorkflowServiceImpl interface.
func (s *WorkflowServiceImpl) GetWorkflowNodeByID(ctx context.Context, req *base.IDReq) (resp *workflowsvc.WorkflowNode, err error) {
	return
}

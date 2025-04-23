package handler

import (
	"context"

	"github.com/aiagt/aiagt/common/types"
	"github.com/aiagt/aiagt/pkg/schema"
	"github.com/aiagt/aiagt/pkg/workflow"

	"github.com/aiagt/aiagt/apps/workflow/model"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc"
)

// CreateWorkflowNode implements the WorkflowServiceImpl interface.
func (s *WorkflowServiceImpl) CreateWorkflowNode(ctx context.Context, req *workflowsvc.CreateWorkflowNodeReq) (resp *base.Empty, err error) {
	wf, err := s.workflowDao.GetByID(ctx, req.GetWorkflowId())
	if err != nil {
		return nil, bizCreateWorkflowNode.NewErr(err).Log(ctx, "get workflow error")
	}

	if ctxutil.Forbidden(ctx, wf.AuthorID) {
		return nil, bizCreateWorkflowNode.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
	}

	workflowNode := model.WorkflowNode{
		Name:       req.GetName(),
		WorkflowID: req.GetWorkflowId(),
		NextIDs:    req.GetNextIds(),
		Type:       model.WorkflowNodeType(req.GetType()),
	}

	workflowNode.InputMapper, err = types.GetValue[workflow.ObjectMapper](req.GetInputMapper())
	if err != nil {
		return nil, bizCreateWorkflowNode.NewErr(err).Log(ctx, "decode input mapper error")
	}

	workflowNode.OutputSchema, err = types.GetValue[*schema.Definition](req.GetOutputSchema())
	if err != nil {
		return nil, bizCreateWorkflowNode.NewErr(err).Log(ctx, "decode output mapper error")
	}

	workflowNode.BatchField, err = types.GetValue[*workflow.ObjectField](req.GetBatchField())
	if err != nil {
		return nil, bizCreateWorkflowNode.NewErr(err).Log(ctx, "decode batch field error")
	}

	workflowNode.NodeParams, err = types.GetValue[model.WorkflowNodeParams](req.GetNodeParams())
	if err != nil {
		return nil, bizCreateWorkflowNode.NewErr(err).Log(ctx, "decode node params error")
	}

	err = s.workflowNodeDao.Create(ctx, &workflowNode)
	if err != nil {
		return nil, bizCreateWorkflowNode.NewErr(err).Log(ctx, "create workflow error")
	}

	return
}

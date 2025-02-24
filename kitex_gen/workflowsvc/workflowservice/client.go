// Code generated by Kitex v0.10.0. DO NOT EDIT.

package workflowservice

import (
	"context"
	workflowsvc "github.com/aiagt/aiagt/kitex_gen/workflowsvc"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CallWorkflow(ctx context.Context, req *workflowsvc.CallWorkflowReq, callOptions ...callopt.Option) (r *workflowsvc.CallWorkflowResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kWorkflowServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kWorkflowServiceClient struct {
	*kClient
}

func (p *kWorkflowServiceClient) CallWorkflow(ctx context.Context, req *workflowsvc.CallWorkflowReq, callOptions ...callopt.Option) (r *workflowsvc.CallWorkflowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CallWorkflow(ctx, req)
}

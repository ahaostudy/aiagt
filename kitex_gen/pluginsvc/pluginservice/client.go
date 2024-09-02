// Code generated by Kitex v0.10.0. DO NOT EDIT.

package pluginservice

import (
	"context"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	pluginsvc "github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Create(ctx context.Context, req *pluginsvc.CreatePluginReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	Update(ctx context.Context, req *pluginsvc.UpdatePluginReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	List(ctx context.Context, req *base.PaginationReq, callOptions ...callopt.Option) (r *pluginsvc.ListPluginResp, err error)
	GetByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *pluginsvc.Plugin, err error)
	Delete(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	CreateTool(ctx context.Context, req *pluginsvc.CreatePluginToolReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	UpdateTool(ctx context.Context, req *pluginsvc.UpdatePluginToolReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	ListTool(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *pluginsvc.ListPluginToolResp, err error)
	GetToolByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *pluginsvc.PluginTool, err error)
	DeleteTool(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	CallPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq, callOptions ...callopt.Option) (r *pluginsvc.CallPluginToolResp, err error)
	TestPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq, callOptions ...callopt.Option) (r *pluginsvc.TestPluginToolResp, err error)
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
	return &kPluginServiceClient{
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

type kPluginServiceClient struct {
	*kClient
}

func (p *kPluginServiceClient) Create(ctx context.Context, req *pluginsvc.CreatePluginReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Create(ctx, req)
}

func (p *kPluginServiceClient) Update(ctx context.Context, req *pluginsvc.UpdatePluginReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Update(ctx, req)
}

func (p *kPluginServiceClient) List(ctx context.Context, req *base.PaginationReq, callOptions ...callopt.Option) (r *pluginsvc.ListPluginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.List(ctx, req)
}

func (p *kPluginServiceClient) GetByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *pluginsvc.Plugin, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetByID(ctx, req)
}

func (p *kPluginServiceClient) Delete(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Delete(ctx, req)
}

func (p *kPluginServiceClient) CreateTool(ctx context.Context, req *pluginsvc.CreatePluginToolReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateTool(ctx, req)
}

func (p *kPluginServiceClient) UpdateTool(ctx context.Context, req *pluginsvc.UpdatePluginToolReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTool(ctx, req)
}

func (p *kPluginServiceClient) ListTool(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *pluginsvc.ListPluginToolResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListTool(ctx, req)
}

func (p *kPluginServiceClient) GetToolByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *pluginsvc.PluginTool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetToolByID(ctx, req)
}

func (p *kPluginServiceClient) DeleteTool(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTool(ctx, req)
}

func (p *kPluginServiceClient) CallPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq, callOptions ...callopt.Option) (r *pluginsvc.CallPluginToolResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CallPluginTool(ctx, req)
}

func (p *kPluginServiceClient) TestPluginTool(ctx context.Context, req *pluginsvc.CallPluginToolReq, callOptions ...callopt.Option) (r *pluginsvc.TestPluginToolResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TestPluginTool(ctx, req)
}

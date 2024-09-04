// Code generated by Kitex v0.10.0. DO NOT EDIT.

package userservice

import (
	"context"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, req *usersvc.RegisterReq, callOptions ...callopt.Option) (r *usersvc.RegisterResp, err error)
	Login(ctx context.Context, req *usersvc.LoginReq, callOptions ...callopt.Option) (r *usersvc.LoginResp, err error)
	ForgotPassword(ctx context.Context, req *usersvc.ForgotPasswordReq, callOptions ...callopt.Option) (r *usersvc.ForgotPasswordResp, err error)
	UpdateUser(ctx context.Context, req *usersvc.UpdateUserReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	GetUserByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *usersvc.User, err error)
	CreateSecret(ctx context.Context, req *usersvc.CreateSecretReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	UpdateSecret(ctx context.Context, req *usersvc.UpdateSecretReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	DeleteSecret(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Empty, err error)
	ListSecret(ctx context.Context, req *usersvc.ListSecretReq, callOptions ...callopt.Option) (r *usersvc.ListSecretResp, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, req *usersvc.RegisterReq, callOptions ...callopt.Option) (r *usersvc.RegisterResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kUserServiceClient) Login(ctx context.Context, req *usersvc.LoginReq, callOptions ...callopt.Option) (r *usersvc.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kUserServiceClient) ForgotPassword(ctx context.Context, req *usersvc.ForgotPasswordReq, callOptions ...callopt.Option) (r *usersvc.ForgotPasswordResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ForgotPassword(ctx, req)
}

func (p *kUserServiceClient) UpdateUser(ctx context.Context, req *usersvc.UpdateUserReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, req)
}

func (p *kUserServiceClient) GetUserByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *usersvc.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserByID(ctx, req)
}

func (p *kUserServiceClient) CreateSecret(ctx context.Context, req *usersvc.CreateSecretReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateSecret(ctx, req)
}

func (p *kUserServiceClient) UpdateSecret(ctx context.Context, req *usersvc.UpdateSecretReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateSecret(ctx, req)
}

func (p *kUserServiceClient) DeleteSecret(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteSecret(ctx, req)
}

func (p *kUserServiceClient) ListSecret(ctx context.Context, req *usersvc.ListSecretReq, callOptions ...callopt.Option) (r *usersvc.ListSecretResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListSecret(ctx, req)
}
package main

import (
	"context"
	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *usersvc.RegisterReq) (resp *usersvc.RegisterResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *usersvc.LoginReq) (resp *usersvc.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// ForgotPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) ForgotPassword(ctx context.Context, req *usersvc.ForgotPasswordReq) (resp *usersvc.ForgotPasswordResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *usersvc.UpdateUserReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// GetUserByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByID(ctx context.Context, req *base.IDReq) (resp *usersvc.User, err error) {
	// TODO: Your code here...
	return
}

// CreateSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateSecret(ctx context.Context, req *usersvc.CreateSecretReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// UpdateSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateSecret(ctx context.Context, req *usersvc.UpdateSecretReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// DeleteSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteSecret(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	// TODO: Your code here...
	return
}

// ListSecret implements the UserServiceImpl interface.
func (s *UserServiceImpl) ListSecret(ctx context.Context, req *usersvc.ListSecretReq) (resp *usersvc.ListSecretResp, err error) {
	// TODO: Your code here...
	return
}
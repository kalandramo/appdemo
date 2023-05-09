package main

import (
	"context"

	"github.com/kalandramo/appdemo/cmd/user/service"
	user "github.com/kalandramo/appdemo/kitex_gen/user"
	"github.com/kalandramo/appdemo/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

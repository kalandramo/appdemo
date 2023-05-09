package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/kalandramo/appdemo/cmd/user/infras/mysql"
	"github.com/kalandramo/appdemo/kitex_gen/user"
	"github.com/kalandramo/appdemo/pkg/errno"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) error {
	users, err := mysql.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return mysql.CreateUser(s.ctx, []*mysql.User{{
		Username: req.Username,
		Password: password,
	}})
}

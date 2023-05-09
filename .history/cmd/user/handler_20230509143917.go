package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/kalandramo/appdemo/cmd/user/infras/mysql"
	user "github.com/kalandramo/appdemo/kitex_gen/user"
	"github.com/kalandramo/appdemo/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	users, err := db.QueryUser(s.ctx, req.Username)
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
	return mysql.CreateUser(s.ctx, []*db.User{{
		Username: req.Username,
		Password: password,
	}})
}

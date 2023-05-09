package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/kalandramo/appdemo/cmd/user/infras/mysql"
	user "github.com/kalandramo/appdemo/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
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

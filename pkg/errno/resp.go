package errno

import (
	"errors"
	"time"

	"github.com/kalandramo/appdemo/kitex_gen/base"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *base.BaseResp {
	if err == nil {
		return baseResp(Success)
	}
	e := ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err ErrNo) *base.BaseResp {
	return &base.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}

package errno

import (
	"errors"
	"fmt"
)

const (
	// System Code
	SuccessCode    = 0
	ServiceErrCode = 10001
	ParamErrCode   = 10002

	// User ErrCode
	UserNotExistErrCode     = 11002
	UserAlreadyExistErrCode = 11003
	AuthorizationFailedCode = 11001
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success    = NewErrNo(int64(SuccessCode), "Success")
	ServiceErr = NewErrNo(int64(ServiceErrCode), "Service is unable to start successfully")
	ParamErr   = NewErrNo(int64(ParamErrCode), "Wrong Parameter has been given")

	UserNotExistErr        = NewErrNo(int64(UserNotExistErrCode), "User does not exists")
	UserAlreadyExistErr    = NewErrNo(int64(UserAlreadyExistErrCode), "User already exists")
	AuthorizationFailedErr = NewErrNo(int64(AuthorizationFailedCode), "Authorization failed")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

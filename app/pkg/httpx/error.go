package httpx

import (
	"github.com/pkg/errors"
	"net/http"
)

var (
	InternalServerError   = errors.New("internal server error")
	MethodNotAllowedError = errors.New("method not allowed")
	ResourceNotFoundError = errors.New("resource not found")
)

type Error struct {
	Code int
	Err  string
}

func (e Error) Error() string {
	if e.Err == "" {
		return ""
	}
	return e.Err
}

func NewErrorMsg(code int, msg string) Error {
	return NewError(code, errors.New(msg))
}

func NewError(code int, err error) Error {
	return Error{
		Code: code,
		Err:  err.Error(),
	}
}

func NotFoundError(err error) Error {
	return NewError(http.StatusNotFound, err)
}

func ForbiddenError(err error) Error {
	return NewError(http.StatusForbidden, err)
}

func UnAuthorizedError(err error) Error {
	return NewError(http.StatusUnauthorized, err)
}

func InternalError(err error) Error {
	return NewError(http.StatusInternalServerError, err)
}

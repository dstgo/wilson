package httputil

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type Error struct {
	code  int
	error string
}

func (e *Error) Error() string {
	if e.error == "" {
		return ""
	}
	return fmt.Sprintf("http Error %d: %s", e.code, e.error)
}

func NewErrorMsg(code int, msg string) *Error {
	return NewError(code, errors.New(msg))
}

func NewError(code int, err error) *Error {
	return &Error{
		code:  code,
		error: err.Error(),
	}
}

func NotFoundError(err error) *Error {
	return NewError(http.StatusNotFound, err)
}

func ForbiddenError(err error) *Error {
	return NewError(http.StatusForbidden, err)
}

func UnAuthorizedError(err error) *Error {
	return NewError(http.StatusUnauthorized, err)
}

func InternalError(err error) *Error {
	return NewError(http.StatusInternalServerError, err)
}

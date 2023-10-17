package errs

import (
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/types/code"
	"net/http"
)

func NewErr() *ResponseError {
	return new(ResponseError)
}

// ResponseError
// a response error wrap
// Er field
type ResponseError struct {
	ErrorCode  int
	HttpStatus int
	LangCode   string
	Er         error
}

func (e *ResponseError) Code(code int) *ResponseError {
	e.ErrorCode = code
	return e
}

func (e *ResponseError) Status(status int) *ResponseError {
	e.HttpStatus = status
	return e
}

func (e *ResponseError) I18n(langCode string) *ResponseError {
	e.LangCode = langCode
	return e
}

func (e *ResponseError) Err(err error) *ResponseError {
	e.Er = err
	return e
}

func (e *ResponseError) Error() string {
	if e.LangCode == "" {
		return e.Er.Error()
	}

	if e.Er != nil {
		return WrapI18n(e.Er, e.LangCode).Error()
	}

	return locale.Get(e.LangCode)
}

// response error helper function

func BadRequest(err error) *ResponseError {
	return &ResponseError{
		HttpStatus: http.StatusBadRequest,
		Er:         err,
	}
}

func UnAuthorized(err error) *ResponseError {
	return &ResponseError{
		HttpStatus: http.StatusUnauthorized,
		Er:         err,
	}
}

func Forbidden(err error) *ResponseError {
	return &ResponseError{
		HttpStatus: http.StatusForbidden,
		Er:         err,
	}
}

func ResourceNotFound(err error) *ResponseError {
	return &ResponseError{
		ErrorCode:  code.ResourceNotFound,
		HttpStatus: http.StatusNotFound,
		Er:         err,
	}
}

func DataBaseErr(err error) *ResponseError {
	return &ResponseError{
		ErrorCode:  code.DatabaseError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.database",
		Er:         err,
	}
}

func FileSystemErr(err error) *ResponseError {
	return &ResponseError{
		ErrorCode:  code.FilesystemError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.filesystem",
		Er:         err,
	}
}

func NetworkErr(err error) *ResponseError {
	return &ResponseError{
		ErrorCode:  code.NetworkError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.network",
		Er:         err,
	}
}

func ProgramErr(err error) *ResponseError {
	return &ResponseError{
		ErrorCode:  code.UnknownError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.program",
		Er:         err,
	}
}

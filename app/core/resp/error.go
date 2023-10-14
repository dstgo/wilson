package resp

import (
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/pkg/errorx"
	"github.com/dstgo/wilson/app/types/code"
	"net/http"
)

func NewErr() *ResponseError {
	return new(ResponseError)
}

// ResponseError
// a response error wrap
// err field
type ResponseError struct {
	CustomCode int
	HttpStatus int
	LangCode   string
	err        error
}

func (e *ResponseError) Code(code int) *ResponseError {
	e.CustomCode = code
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
	e.err = err
	return e
}

func (e *ResponseError) Error() string {
	if e.LangCode == "" {
		return e.err.Error()
	}

	if e.err != nil {
		return errorx.WrapI18n(e.err, e.LangCode).Error()
	}

	return locale.Get(e.LangCode)
}

// helper function

func DataBaseErr(err error) *ResponseError {
	return &ResponseError{
		CustomCode: code.DatabaseError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.database",
		err:        err,
	}
}

func FileSystemErr(err error) *ResponseError {
	return &ResponseError{
		CustomCode: code.FilesystemError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.filesystem",
		err:        err,
	}
}

func NetworkErr(err error) *ResponseError {
	return &ResponseError{
		CustomCode: code.NetworkError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.network",
		err:        err,
	}
}

func ProgramErr(err error) *ResponseError {
	return &ResponseError{
		CustomCode: code.UnknownError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.program",
		err:        err,
	}
}

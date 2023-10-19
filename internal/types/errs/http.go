package errs

import (
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/types/code"
	"net/http"
)

func NewError() *LocaleError {
	return new(LocaleError)
}

func NewI18nError(i18n string) *LocaleError {
	return NewError().I18n(i18n)
}

// LocaleError
// a response error wrap
// Er field
type LocaleError struct {
	ErrorCode  int
	HttpStatus int
	LangCode   string
	Er         error
}

func (e *LocaleError) Code(code int) *LocaleError {
	e.ErrorCode = code
	return e
}

func (e *LocaleError) Status(status int) *LocaleError {
	e.HttpStatus = status
	return e
}

func (e *LocaleError) I18n(langCode string) *LocaleError {
	e.LangCode = langCode
	return e
}

func (e *LocaleError) Err(err error) *LocaleError {
	e.Er = err
	return e
}

func (e *LocaleError) Error() string {
	if e.LangCode == "" {
		return e.Er.Error()
	}

	if e.Er != nil {
		return WrapI18n(e.Er, e.LangCode).Error()
	}

	return locale.Get(e.LangCode)
}

// response error helper function

func BadRequest(err error) *LocaleError {
	return &LocaleError{
		HttpStatus: http.StatusBadRequest,
		Er:         err,
	}
}

func UnAuthorized(err error) *LocaleError {
	return &LocaleError{
		HttpStatus: http.StatusUnauthorized,
		Er:         err,
	}
}

func Forbidden(err error) *LocaleError {
	return &LocaleError{
		HttpStatus: http.StatusForbidden,
		Er:         err,
	}
}

func ResourceNotFound(err error) *LocaleError {
	return &LocaleError{
		ErrorCode:  code.ResourceNotFound,
		HttpStatus: http.StatusNotFound,
		Er:         err,
	}
}

func DataBaseErr(err error) *LocaleError {
	return &LocaleError{
		ErrorCode:  code.DatabaseError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.database",
		Er:         err,
	}
}

func FileSystemErr(err error) *LocaleError {
	return &LocaleError{
		ErrorCode:  code.FilesystemError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.filesystem",
		Er:         err,
	}
}

func NetworkErr(err error) *LocaleError {
	return &LocaleError{
		ErrorCode:  code.NetworkError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.network",
		Er:         err,
	}
}

func ProgramErr(err error) *LocaleError {
	return &LocaleError{
		ErrorCode:  code.UnknownError,
		HttpStatus: http.StatusInternalServerError,
		LangCode:   "err.program",
		Er:         err,
	}
}

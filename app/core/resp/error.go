package resp

import (
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	// http status code will be response
	Code int
	// language code
	LangCode string
	// error message
	Err string
}

func (e Error) Error() string {
	return e.Err
}

func (e Error) ErrorCtx(ctx *gin.Context) string {
	return locale.L().GetWithCtx(ctx, e.LangCode)
}

func NewI18nErr(code int, langcode string) Error {
	return Error{
		Code:     code,
		LangCode: langcode,
	}
}

func NewMsgErr(code int, msg string) Error {
	return Error{
		Code: code,
		Err:  msg,
	}
}

func NewErr(code int, err error) Error {
	return Error{
		Code: code,
		Err:  err.Error(),
	}
}

func NotFoundError(err error) Error {
	return NewErr(http.StatusNotFound, err)
}

func ForbiddenError(err error) Error {
	return NewErr(http.StatusForbidden, err)
}

func UnAuthorizedError(err error) Error {
	return NewErr(http.StatusUnauthorized, err)
}

func InternalError(err error) Error {
	return NewErr(http.StatusInternalServerError, err)
}

package resp

import (
	"errors"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	// custom Code
	Code int `json:"code"`
	// response mgs
	Msg string `json:"msg,omitempty"`
	// response Data
	Data any `json:"data,omitempty"`
	// response Error
	Error string `json:"error,omitempty"`
}

func Resp(ctx *gin.Context, httpCode int, code int, msg string, data any, err error) {
	body := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	if err != nil {
		var e Error
		if errors.As(err, &e) {
			httpCode = e.Code
		}
		if len(e.LangCode) > 0 {
			body.Error = e.ErrorCtx(ctx)
		} else {
			body.Error = err.Error()
		}
	}

	ctx.JSON(httpCode, body)
}

func Ok(ctx *gin.Context, code int, msg string, data any) {
	Resp(ctx, http.StatusOK, code, msg, data, nil)
}

func OkI18n(ctx *gin.Context, code int, langCode string, data any) {
	Resp(ctx, http.StatusOK, code, locale.L().GetWithCtx(ctx, langCode), data, nil)
}

func OkNilBody(ctx *gin.Context, code int, msg string) {
	Resp(ctx, http.StatusOK, code, msg, nil, nil)
}

func OkNilBodyI18n(ctx *gin.Context, code int, langCode string) {
	Resp(ctx, http.StatusOK, code, locale.L().GetWithCtx(ctx, langCode), nil, nil)
}

func Fail(ctx *gin.Context, code int, err error) {
	Resp(ctx, http.StatusBadRequest, code, "", nil, err)
}

func FailBody(ctx *gin.Context, code int, body any, err error) {
	Resp(ctx, http.StatusBadRequest, code, "", body, err)
}

func NotFound(ctx *gin.Context, code int, err error) {
	Resp(ctx, http.StatusNotFound, code, "", nil, err)
}

func Forbidden(ctx *gin.Context, code int, err error) {
	Resp(ctx, http.StatusForbidden, code, "", nil, err)
}

func MethodNotAllowed(ctx *gin.Context, code int, err error) {
	Resp(ctx, http.StatusMethodNotAllowed, code, "", nil, err)
}

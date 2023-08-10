package httpx

import (
	"errors"
	"github.com/gin-gonic/gin"
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

func Resp(ctx *gin.Context, code int, msg string, data any, err error) {
	var httpCode int
	body := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	httpCode = 200
	if err != nil {
		httpCode = 400
		var e *Error
		if errors.As(err, &e) {
			httpCode = e.code
		}
		body.Error = err.Error()
	}
	ctx.JSON(httpCode, body)
}

func Ok(ctx *gin.Context, code int, msg string, data any) {
	Resp(ctx, code, msg, data, nil)
}

func Failed(ctx *gin.Context, code int, err error) {
	Resp(ctx, code, "", nil, err)
}

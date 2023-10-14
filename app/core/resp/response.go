package resp

import (
	"errors"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx: ctx}
}

// Ok means that the request is successful.
func Ok(ctx *gin.Context) *Response {
	return NewResponse(ctx).Status(http.StatusOK)
}

// Fail means that the request is unsuccessful, the result usually caused by client, not server
// like incorrect parameters
func Fail(ctx *gin.Context) *Response {
	return NewResponse(ctx).Status(http.StatusBadRequest)
}

// Error means that the request is unsuccessful, the reason usually caused by server
func Error(ctx *gin.Context) *Response {
	return NewResponse(ctx).Status(http.StatusInternalServerError)
}

type Response struct {
	err    error
	status int
	ctx    *gin.Context

	// custom CustomCode
	CustomCode int `json:"code"`
	// response mgs
	Message string `json:"msg,omitempty"`
	// response Payload
	Payload any `json:"data,omitempty"`
	// response ErrorMsg
	ErrorMsg string `json:"error,omitempty"`
}

func (r *Response) Status(status int) *Response {
	r.status = status
	return r
}
func (r *Response) Code(code int) *Response {
	r.CustomCode = code
	return r
}

func (r *Response) MsgI18n(langCode string) *Response {
	if r.ctx != nil {
		return r.Msg(locale.GetWithCtx(r.ctx, langCode))
	}
	return r
}

func (r *Response) Msg(msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) Data(data any) *Response {
	r.Payload = data
	return r
}

func (r *Response) Error(err error) *Response {
	r.err = err
	return r
}

func (r *Response) Send() {
	if r.ctx != nil {
		if r.err != nil {
			r.ErrorMsg = r.err.Error()

			var e *ResponseError
			if errors.As(r.err, &e) {

				// if httpcode >= 500, which means internal server error happened
				// for non-internal errors, detailed error information can be displayed externally
				// otherwise only simple description information should be returned to avoid leaking sensitive data
				if e.HttpStatus >= 500 {
					r.ErrorMsg = locale.GetWithCtx(r.ctx, e.LangCode)
				} else {
					r.ErrorMsg = e.Error()
				}

				if e.CustomCode > 0 {
					r.CustomCode = e.CustomCode
				}
			}

			if len(r.ErrorMsg) == 0 {
				r.ErrorMsg = locale.GetWithCtx(r.ctx, "err.unknown")
			}

			r.ctx.Error(r.err)
		}
		r.ctx.JSON(r.status, r)
	}
}

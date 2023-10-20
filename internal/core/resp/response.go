package resp

import (
	"errors"
	"github.com/dstgo/wilson/internal/pkg/httpx"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewResponse(ctx *gin.Context) *Response {
	resp := &Response{ctx: ctx}
	// get accept languages
	resp.locale = httpx.GetFirstAcceptLanguage(ctx)
	return resp
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

// InternalFailed means that the request is unsuccessful, the reason usually caused by server
func InternalFailed(ctx *gin.Context) *Response {
	return NewResponse(ctx).Status(http.StatusInternalServerError)
}

type Response struct {
	// current response error
	err error
	// http status
	status int
	// current context language
	locale string
	ctx    *gin.Context

	// i18n message code
	i18n string
	// fallback message
	fallback string

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
	r.i18n = langCode
	return r
}

func (r *Response) Msg(msg string) *Response {
	r.fallback = msg
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
	if r.ctx == nil {
		panic("response gin context is nil")
	}

	if r.err != nil {
		r.ErrorMsg = r.err.Error()

		var e *errs.LocaleError
		if errors.As(r.err, &e) {
			errMsg := locale.GetWithLang(r.locale, e.LangCode)
			// if httpcode >= 500, which means internal server error happened.
			// for non-internal errors, detailed error information can be displayed externally
			// otherwise only simple description information can be returned to avoid leaking sensitive data
			if e.HttpStatus < 500 && e.Er != nil {
				errMsg = errs.Wrap(e.Er, errMsg).Error()
			}

			// overwrite http status
			if e.HttpStatus > 0 {
				r.status = e.HttpStatus
			}

			// overwrite custom code
			if e.ErrorCode > 0 {
				r.CustomCode = e.ErrorCode
			}

			// error msg fallback
			if len(errMsg) == 0 {
				errMsg = e.LangCode
			}

			r.ErrorMsg = errMsg
		}

		if len(r.ErrorMsg) == 0 {
			r.ErrorMsg = locale.GetWithLang(r.locale, "err.unknown")
		}

		r.ctx.Error(r.err)
	}

	if r.CustomCode == 0 {
		r.CustomCode = r.status * 10
	}

	// try to get localized message
	if len(r.i18n) > 0 {
		r.Message = locale.GetWithLang(r.locale, r.i18n)
		// i18n fallback
		if len(r.Message) == 0 {
			r.Message = r.i18n
		}
	}

	// fallback default message
	if len(r.Message) == 0 {
		r.Message = r.fallback
	}

	r.ctx.JSON(r.status, r)
}

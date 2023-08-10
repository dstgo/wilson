package middleware

import (
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/dstgo/wilson/app/pkg/locale"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func UseCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UseAcceptLanguage(defaultLanguage string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		acceptLang := defaultLanguage
		// choose the one with the highest weight among candidate languages
		candidates := httpx.GetAcceptLanguage(ctx)
		if len(candidates) > 0 {
			acceptLang = candidates[0]
		}
		ctx.Next()
		ctx.Writer.Header().Set(httpx.ContentLanguage, acceptLang)
	}
}

// NotFoundHandler
// return gin.HandlerFunc
func NotFoundHandler(l *locale.Locale) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpx.Failed(ctx, 4040, httpx.NotFoundError(errors.New(l.GetWithCtx(ctx, "http.404"))))
	}
}

func NoMethodHandler(l *locale.Locale) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpx.Failed(ctx, 4050, httpx.NotFoundError(errors.New(l.GetWithCtx(ctx, "http.405"))))
	}
}

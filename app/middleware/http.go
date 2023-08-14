package middleware

import (
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/gin-gonic/gin"
)

func UseCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// UseAcceptLanguage get accept language from http header
func UseAcceptLanguage(defaultLanguage string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		acceptLang := defaultLanguage
		// choose the one with the highest weight among candidate languages
		candidates := httpx.GetAcceptLanguage(ctx)
		if len(candidates) > 0 {
			acceptLang = candidates[0]
		}
		ctx.Next()
		ctx.Writer.Header().Set(httpx.ContentLanguageHeader, acceptLang)
	}
}

// NotFoundHandler
// param l *locale.Locale
// return gin.HandlerFunc
func NotFoundHandler(l *locale.Locale) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpx.Failed(ctx, 4040, httpx.NotFoundError(l.NewErrorWithCtx(ctx, "http.404")))
	}
}

// NoMethodHandler
// param l *locale.Locale
// return gin.HandlerFunc
func NoMethodHandler(l *locale.Locale) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpx.Failed(ctx, 4050, httpx.NotFoundError(l.NewErrorWithCtx(ctx, "http.405")))
	}
}

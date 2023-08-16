package middleware

import (
	"errors"
	"fmt"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/dstgo/wilson/app/pkg/httpx/httpheader"
	"github.com/dstgo/wilson/app/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

// UseCors gin cors middleware
func UseCors(cors *httpx.Cors) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if cors == nil || !cors.Enabled {
			ctx.Next()
			return
		}
		origin := ctx.GetHeader(httpheader.Origin)

		corsHeader := ctx.Writer.Header()
		if len(cors.AllowOrigins) == 1 && cors.AllowOrigins[0] == "*" {
			corsHeader.Set(httpheader.AccessControlAllowOrigin, "*")
		} else if slices.Contains(cors.AllowOrigins, origin) {
			corsHeader.Set(httpheader.AccessControlAllowOrigin, origin)
		}
		corsHeader.Set(httpheader.AccessControlAllowMethods, cors.AccessAllowMethods())
		corsHeader.Set(httpheader.AccessControlAllowHeaders, cors.AccessAllowHeaders())
		corsHeader.Set(httpheader.AccessControlExposeHeaders, cors.AccessExposedHeaders())
		corsHeader.Set(httpheader.AccessControlMaxAge, cors.AccessMaxAge())
		corsHeader.Set(httpheader.AccessControlAllowCredentials, cors.AccessCredentials())

		if ctx.Request.Method == http.MethodOptions {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// UseRecovery handler chain recover
func UseRecovery(logger *logrus.Logger, lang *locale.Locale) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			var (
				responseErr = httpx.InternalServerError
				brokenPipe  bool
			)

			var err any

			if panicErr := recover(); panicErr != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				if ne, ok := panicErr.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne.Err, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				err = panicErr

				request, _ := httputil.DumpRequest(ctx.Request, false)

				entry := logger.
					WithField(types.LogRecoverRequestKey, string(request)).
					WithField(types.LogRecoverErrorKey, fmt.Sprintf("%v", err)).
					WithField(types.LogRequestIdKey, httpx.GetRequestId(ctx))

				if !brokenPipe {
					entry = entry.WithField("stack", string(debug.Stack()))
				}

				if logger != nil {
					entry.Errorln()
				}

				if lang != nil {
					responseErr = lang.NewErrorWithCtx(ctx, "http.500")
				}

				ctx.Abort()

				if brokenPipe {
					return
				}

				httpx.Failed(ctx, 5000, httpx.NewError(http.StatusInternalServerError, responseErr))
			}
		}()

		ctx.Next()
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
		ctx.Writer.Header().Set(httpheader.ContentLanguage, acceptLang)
		ctx.Next()
	}
}

// NotFoundHandler 404 handler
func NotFoundHandler(l *locale.Locale) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := httpx.ResourceNotFoundError
		if l != nil {
			err = l.NewErrorWithCtx(ctx, "http.404")
		}
		httpx.Failed(ctx, 4040, httpx.NotFoundError(err))
	}
}

// NoMethodHandler no method handler
func NoMethodHandler(l *locale.Locale) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := httpx.MethodNotAllowedError
		if l != nil {
			err = l.NewErrorWithCtx(ctx, "http.405")
		}
		httpx.Failed(ctx, 4050, httpx.NotFoundError(err))
	}
}

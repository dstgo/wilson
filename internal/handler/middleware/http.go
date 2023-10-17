package middleware

import (
	"errors"
	"fmt"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/pkg/httpx"
	"github.com/dstgo/wilson/internal/pkg/httpx/httpheader"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/code"
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
func UseRecovery(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {

			if panicErr := recover(); panicErr != nil {
				var (
					brokenPipe bool
					readBody   bool
				)

				var err any
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

				contentType := httpx.GetContentType(ctx)
				if slices.Contains([]string{"application/json"}, contentType) {
					readBody = true
				}

				request, _ := httputil.DumpRequest(ctx.Request, readBody)

				entry := logger.
					WithField(types.LogRecoverRequestKey, string(request)).
					WithField(types.LogRecoverErrorKey, fmt.Sprintf("%v", err)).
					WithField(types.LogRequestIdKey, httpx.GetRequestId(ctx))

				if !brokenPipe {
					entry = entry.WithField(types.LogRecoverStackKey, string(debug.Stack()))
				}

				if logger != nil {
					entry.Errorln()
				}

				ctx.Abort()

				if brokenPipe {
					return
				}

				resp.InternalFailed(ctx).Code(code.InternalServerError).MsgI18n("err.internal").Send()
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
func NotFoundHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp.Fail(ctx).Status(http.StatusNotFound).Code(4040).MsgI18n("err.notfound").Send()
	}
}

// NoMethodHandler no method handler
func NoMethodHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp.Fail(ctx).Status(http.StatusMethodNotAllowed).Code(4050).MsgI18n("err.methodnotallowed").Send()
	}
}

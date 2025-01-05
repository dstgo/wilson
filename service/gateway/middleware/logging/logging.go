package logging

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/dstgo/wilson/service/gateway/config"
	middleware2 "github.com/dstgo/wilson/service/gateway/middleware"
)

func init() {
	middleware2.Register("logging", Middleware)
}

// Middleware is a logging middleware.
func Middleware(c *config.Middleware) (middleware2.Middleware, error) {
	return func(next http.RoundTripper) http.RoundTripper {
		return middleware2.RoundTripperFunc(func(req *http.Request) (reply *http.Response, err error) {
			startTime := time.Now()
			reply, err = next.RoundTrip(req)
			level := log.LevelInfo
			code := http.StatusBadGateway
			errMsg := ""
			if err != nil {
				level = log.LevelError
				errMsg = err.Error()
			} else {
				code = reply.StatusCode
			}
			ctx := req.Context()
			// nodes, _ := middleware.RequestBackendsFromContext(ctx)
			reqOpt, _ := middleware2.FromRequestContext(ctx)
			log.Context(ctx).Log(level,
				"source", "accesslog",
				"host", req.Host,
				"method", req.Method,
				"scheme", req.URL.Scheme,
				"path", req.URL.Path,
				"query", req.URL.RawQuery,
				"code", code,
				"error", errMsg,
				"latency", time.Since(startTime).Seconds(),
				"backend", strings.Join(reqOpt.Backends, ","),
				"backend_code", reqOpt.UpstreamStatusCode,
				"backend_latency", reqOpt.UpstreamResponseTime,
				"last_attempt", reqOpt.LastAttempt,
				"trace", tracing.TraceID()(ctx),
				"span", tracing.SpanID()(ctx),
			)
			return reply, err
		})
	}, nil
}

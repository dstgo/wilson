package middleware

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/peer"

	"github.com/dstgo/wilson/framework/kratosx/library/ip"
)

func IP() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			ipAddr := ""
			if p, is := peer.FromContext(ctx); is { //grpc
				if strings.Contains(p.Addr.String(), "::1") {
					ipAddr = "localhost"
				} else {
					ipAddr = strings.Split(p.Addr.String(), ":")[0]
				}
			}
			if h, is := http.RequestFromServerContext(ctx); is {
				if strings.Contains(h.RemoteAddr, "::1") {
					ipAddr = "localhost"
				} else {
					ipAddr = strings.Split(h.RemoteAddr, ":")[0]
				}
				if h.Header.Get("x-real-ip") != "" {
					ipAddr = h.Header.Get("x-real-ip")
				}
			}
			ctx = context.WithValue(ctx, ip.CtxKey{}, ipAddr)
			return handler(ctx, req)
		}
	}
}

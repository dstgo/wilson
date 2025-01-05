package bbr

import (
	"bytes"
	"io"
	"net/http"

	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"

	"github.com/dstgo/wilson/service/gateway/config"
	middleware2 "github.com/dstgo/wilson/service/gateway/middleware"
)

var _nopBody = io.NopCloser(&bytes.Buffer{})

func init() {
	middleware2.Register("bbr", Middleware)
}

func Middleware(c *config.Middleware) (middleware2.Middleware, error) {
	limiter := bbr.NewLimiter() //use default settings
	return func(next http.RoundTripper) http.RoundTripper {
		return middleware2.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			done, err := limiter.Allow()
			if err != nil {
				return &http.Response{
					Status:     http.StatusText(http.StatusTooManyRequests),
					StatusCode: http.StatusTooManyRequests,
					Body:       _nopBody,
					Header:     make(http.Header),
				}, nil
			}
			resp, err := next.RoundTrip(req)
			done(ratelimit.DoneInfo{Err: err})
			return resp, err
		})
	}, nil
}

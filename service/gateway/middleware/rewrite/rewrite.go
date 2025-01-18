package rewrite

import (
	"net/http"
	"path"
	"strings"

	"github.com/dstgo/wilson/service/gateway/config"
	gtmiddleware "github.com/dstgo/wilson/service/gateway/middleware"
	"github.com/dstgo/wilson/service/gateway/utils"
)

func init() {
	gtmiddleware.Register("rewrite", Middleware)
}

func stripPrefix(origin string, prefix string) string {
	out := strings.TrimPrefix(origin, prefix)
	if out == "" {
		return "/"
	}
	if out[0] != '/' {
		return path.Join("/", out)
	}
	return out
}

func Middleware(c *config.Middleware) (gtmiddleware.Middleware, error) {
	options := &config.Rewrite{}
	if c.Options != nil {
		if err := utils.Copy(c.Options, options); err != nil {
			return nil, err
		}
	}
	requestHeadersRewrite := options.RequestHeadersRewrite
	responseHeadersRewrite := options.ResponseHeadersRewrite
	return func(next http.RoundTripper) http.RoundTripper {
		return gtmiddleware.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			if options.PathRewrite != "" {
				req.URL.Path = options.PathRewrite
			}
			if options.HostRewrite != "" {
				req.Host = options.HostRewrite
			}
			if options.StripPrefix != "" {
				req.URL.Path = stripPrefix(req.URL.Path, options.StripPrefix)
			}
			if requestHeadersRewrite != nil {
				for key, value := range requestHeadersRewrite.Set {
					req.Header.Set(key, value)
				}
				for key, value := range requestHeadersRewrite.Add {
					req.Header.Add(key, value)
				}
				for _, value := range requestHeadersRewrite.Remove {
					req.Header.Del(value)

				}
			}
			resp, err := next.RoundTrip(req)
			if err != nil {
				return nil, err
			}
			if responseHeadersRewrite != nil {
				for key, value := range responseHeadersRewrite.Set {
					resp.Header.Set(key, value)
				}
				for key, value := range responseHeadersRewrite.Add {
					resp.Header.Add(key, value)
				}
				for _, value := range responseHeadersRewrite.Remove {
					resp.Header.Del(value)

				}
			}
			return resp, nil
		})
	}, nil
}

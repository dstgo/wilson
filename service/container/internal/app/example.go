package app

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/service/container/internal/conf"

	example "github.com/dstgo/wilson/api/gen/container/example/v1"
)

type Example struct {
	example.UnimplementedExampleServer
}

func NewExample(c *conf.Config) *Example {
	return &Example{}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewExample(c)
		example.RegisterExampleServer(gs, srv)
		example.RegisterExampleHTTPServer(hs, srv)
	})
}

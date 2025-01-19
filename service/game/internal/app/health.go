package app

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	health "github.com/dstgo/wilson/api/gen/game/health/v1"
	"github.com/dstgo/wilson/service/game/internal/conf"
)

type Health struct {
	health.UnimplementedHealthServer
}

func NewHealth(c *conf.Config) *Health {
	return &Health{}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewHealth(c)
		health.RegisterHealthServer(gs, srv)
		health.RegisterHealthHTTPServer(hs, srv)
	})
}

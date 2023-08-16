package system

import (
	"github.com/dstgo/wilson/app/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

var SystemApiSet = wire.NewSet(
	NewPingApi,
	NewSystemRouter,
)

type Router struct {
	PingApi *PingApi
}

func NewSystemRouter(root *route.Router, pingApi *PingApi) *Router {
	root.GET("/ping", route.Metas(meta.NoAuth), pingApi.Ping)

	return &Router{PingApi: pingApi}
}

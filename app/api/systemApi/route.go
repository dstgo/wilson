package systemApi

import (
	"github.com/dstgo/wilson/pkg/coco/route"
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
	root.GET("/ping", route.Metas(), pingApi.Ping)
	return &Router{PingApi: pingApi}
}

package systemApi

import (
	route2 "github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

var SystemApiSet = wire.NewSet(
	NewPingApi,
	NewSystemRouter,
)

type Router struct {
	PingApi *PingApi
}

func NewSystemRouter(root *route2.Router, pingApi *PingApi) *Router {
	root.GET("/ping", route2.Metas(), pingApi.Ping)

	return &Router{PingApi: pingApi}
}

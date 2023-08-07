package systemApi

import (
	"github.com/dstgo/wilson/pkg/coco"
	"github.com/google/wire"
)

var SystemApiSet = wire.NewSet(
	NewPingApi,
	NewSystemRouter,
)

type Router struct {
	PingApi *PingApi
}

func NewSystemRouter(coco *coco.Core, ping *PingApi) *Router {
	return &Router{PingApi: ping}
}

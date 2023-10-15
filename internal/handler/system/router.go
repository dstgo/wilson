package system

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

// HandlerRouter just for wire injection, no real influence
type HandlerRouter types.NopType

var SystemRouterSet = wire.NewSet(
	SystemProviderSet,
	wire.Struct(new(Handler), "*"),
	SetupRouter,
)

type Handler struct {
	Ping PingHandler
}

func SetupRouter(api *route.Router, handler Handler) HandlerRouter {
	// system
	{
		api.GET("/ping", route.MetaSum(meta.NoAuth), handler.Ping.Ping)
	}
	return types.NopObj
}

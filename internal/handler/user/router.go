package user

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

// HandlerRouter just for wire injection, no real influence
type HandlerRouter types.NopType

var UserRouterSet = wire.NewSet(
	UserProviderSet,
	wire.Struct(new(Handler), "*"),
	SetupRouter,
)

type Handler struct {
	Info InfoHandler
}

func SetupRouter(api *route.Router, handler Handler) HandlerRouter {

	return types.NopObj
}

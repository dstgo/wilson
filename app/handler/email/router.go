package email

import (
	"github.com/dstgo/wilson/app/types"
	"github.com/dstgo/wilson/app/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

// HandlerRouter just for wire injection, no real influence
type HandlerRouter types.NopType

var EmailRouterSet = wire.NewSet(
	EmailProviderSet,
	wire.Struct(new(Handler), "*"),
	SetupRouter,
)

type Handler struct {
	Email EmailHandler
}

func SetupRouter(api *route.Router, Handler Handler) HandlerRouter {
	emailGroup := api.Group("email", nil)
	{
		emailGroup.GET("code", route.MetaSum(meta.NoAuth), Handler.Email.SendCodeEmail)
	}
	return types.NopObj
}

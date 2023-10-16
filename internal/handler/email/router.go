package email

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/meta"
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
		emailGroup.GET("code", route.Metas(meta.NoAuth, meta.Name("auth.login")), Handler.Email.SendCodeEmail)
	}
	return types.NopObj
}

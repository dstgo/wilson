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
	Auth AuthHandler
}

func SetupRouter(api *route.Router, handler Handler) HandlerRouter {
	// system
	{
		api.GET("/ping", route.Metas(meta.NoAuth, meta.Name("route.sys.ping")), handler.Ping.Ping)
	}
	// auth api
	{
		authGroup := api.Group("auth", nil)
		// POST
		authGroup.POST("/login", route.Metas(meta.NoAuth, meta.Name("route.auth.login")), handler.Auth.Login)
		authGroup.POST("/register", route.Metas(meta.NoAuth, meta.Name("route.auth.register")), handler.Auth.Register)
		authGroup.POST("/forgotpwd", route.Metas(meta.NoAuth, meta.Name("route.auth.forgotpasswd")), handler.Auth.ForgotPassword)
		// DELETE
		authGroup.DELETE("/logout", route.Metas(meta.Name("route.auth.logout")), handler.Auth.Logout)
	}
	return types.NopObj
}

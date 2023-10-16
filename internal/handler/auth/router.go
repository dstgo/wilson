package auth

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

// HandlerRouter just for wire injection, no real influence
type HandlerRouter types.NopType

var AuthRouterSet = wire.NewSet(
	AuthProviderSet,
	SetupRouter,
	wire.Struct(new(Handler), "*"),
)

type Handler struct {
	Auth AuthHandler
}

func SetupRouter(api *route.Router, handler Handler) HandlerRouter {
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

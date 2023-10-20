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
	Role RoleHandler
}

func SetupRouter(api *route.Router, handler Handler) HandlerRouter {
	// system
	{
		api.GET("/ping", route.Metas(meta.NoAuth, meta.Name("route.sys.ping")), handler.Ping.Ping)
		api.GET("/pong", route.Metas(meta.NoAuth, meta.Name("route.sys.ping")), handler.Ping.Pong)
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
	// role api
	roleGroup := api.Group("role", nil)
	{
		roleHandler := handler.Role
		// GET
		roleGroup.GET("list", nil, roleHandler.GetRoleList)
		roleGroup.GET("perms", nil, roleHandler.GetRolePerms)

		// POST
		roleGroup.POST("create", nil, roleHandler.CreateRole)
		roleGroup.POST("update", nil, roleHandler.UpdateRole)
		roleGroup.POST("grant", nil, roleHandler.GrantRolePerms)

		// DELETE
		roleGroup.DELETE("remove", nil, roleHandler.RemoveRole)

		// permission

		// GET
		roleGroup.GET("list", nil, roleHandler.GetPermList)
		// POST
		roleGroup.POST("create", nil, roleHandler.CreatePermission)
		roleGroup.POST("update", nil, roleHandler.UpdatePermission)
		// DELETE
		roleGroup.DELETE("remove", nil, roleHandler.RemovePermission)
	}
	return types.NopObj
}

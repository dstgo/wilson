package system

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/meta"
	"github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/pkg/ginx"
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

func SetupRouter(api *ginx.RouterGroup, handler Handler) HandlerRouter {
	// system
	{
		api.GET("/ping", ginx.M(meta.NoAuth, meta.Name("route.sys.ping")), handler.Ping.Ping)
		api.GET("/pong", ginx.M(meta.NoAuth, meta.Name("route.sys.pong")), handler.Ping.Pong)
	}
	// auth api
	{
		authGroup := api.Group("auth", ginx.M(meta.Name("route.auth")))
		// POST
		authGroup.POST("/login", ginx.M(meta.NoAuth, meta.Name("route.auth.login")), handler.Auth.Login)
		authGroup.POST("/register", ginx.M(meta.NoAuth, meta.Name("route.auth.register")), handler.Auth.Register)
		authGroup.POST("/forgotpwd", ginx.M(meta.NoAuth, meta.Name("route.auth.forgotPasswd")), handler.Auth.ForgotPassword)
		// DELETE
		authGroup.DELETE("/logout", ginx.M(meta.Name("route.auth.logout")), handler.Auth.Logout)
	}
	// role api
	roleGroup := api.Group("/role", ginx.M(meta.Name("route.role"), meta.Roles(role.AdminRole.Code)))
	roleHandler := handler.Role
	{

		// GET
		roleGroup.GET("/list", ginx.M(meta.Name("route.role.list")), roleHandler.GetRoleList)
		roleGroup.GET("/perms", ginx.M(meta.Name("route.role.perms")), roleHandler.GetRolePerms)

		// POST
		roleGroup.POST("/create", ginx.M(meta.Name("route.role.create")), roleHandler.CreateRole)
		roleGroup.POST("/update", ginx.M(meta.Name("route.role.update")), roleHandler.UpdateRole)
		roleGroup.POST("/grant", ginx.M(meta.Name("route.role.grant")), roleHandler.GrantRolePerms)

		// DELETE
		roleGroup.DELETE("/remove", ginx.M(meta.Name("route.role.remove")), roleHandler.RemoveRole)
	}
	// perm api
	permGroup := api.Group("/perm", ginx.M(meta.Name("route.perm"), meta.Roles(role.AdminRole.Code)))
	{
		// GET
		permGroup.GET("/list", ginx.M(meta.Name("route.perm.list")), roleHandler.GetPermList)
		// POST
		permGroup.POST("/create", ginx.M(meta.Name("route.perm.create")), roleHandler.CreatePermission)
		permGroup.POST("/update", ginx.M(meta.Name("route.perm.update")), roleHandler.UpdatePermission)
		// DELETE
		permGroup.DELETE("/remove", ginx.M(meta.Name("route.perm.delete")), roleHandler.RemovePermission)
	}
	return types.NopObj
}

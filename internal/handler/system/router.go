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
	Key  APIKeyHandler
	Dict DictHandler
}

func SetupRouter(api *ginx.RouterGroup, handler Handler) HandlerRouter {
	// system
	systemGroup := api.Group("", ginx.M(meta.Group("route.sys.group")))
	{
		systemGroup.GET("/ping", ginx.M(meta.NoAuth, meta.Name("route.sys.ping")), handler.Ping.Ping)
		systemGroup.GET("/pong", ginx.M(meta.Anonymous, meta.Name("route.sys.pong")), handler.Ping.Pong)
	}

	// auth api
	authGroup := api.Group("auth", ginx.M(meta.Group("route.auth.group")))
	{
		// GET
		authGroup.GET("/refresh", ginx.M(meta.NoAuth, meta.Name("route.auth.refresh")), handler.Auth.Refresh)
		// POST
		authGroup.POST("/login", ginx.M(meta.NoAuth, meta.Name("route.auth.login")), handler.Auth.Login)
		authGroup.POST("/register", ginx.M(meta.NoAuth, meta.Name("route.auth.register")), handler.Auth.Register)
		authGroup.POST("/forgotpwd", ginx.M(meta.NoAuth, meta.Name("route.auth.forgotPasswd")), handler.Auth.ForgotPassword)
		// DELETE
		authGroup.DELETE("/logout", ginx.M(meta.Name("route.auth.logout"), meta.Roles(role.UserRole)), handler.Auth.Logout)
	}

	// role api
	roleGroup := api.Group("/role", ginx.M(meta.Group("route.role.group"), meta.Roles(role.AdminRole)))
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
	permGroup := api.Group("/perm", ginx.M(meta.Group("route.perm.group"), meta.Roles(role.AdminRole)))
	{
		// GET
		permGroup.GET("/list", ginx.M(meta.Name("route.perm.list")), roleHandler.GetPermList)
		// POST
		permGroup.POST("/create", ginx.M(meta.Name("route.perm.create")), roleHandler.CreatePermission)
		permGroup.POST("/update", ginx.M(meta.Name("route.perm.update")), roleHandler.UpdatePermission)
		// DELETE
		permGroup.DELETE("/remove", ginx.M(meta.Name("route.perm.delete")), roleHandler.RemovePermission)
	}

	// api key api
	keyGroup := api.Group("/key", ginx.M(meta.Group("route.key.group"), meta.Roles(role.AdminRole)))
	keyHandler := handler.Key
	{
		// GET
		keyGroup.GET("/list", ginx.M(meta.Name("route.key.list")), keyHandler.ListAPIKeys)
		// POST
		keyGroup.POST("/create", ginx.M(meta.Name("route.key.create")), keyHandler.CreateAPIKey)
		// DELETE
		keyGroup.DELETE("/remove", ginx.M(meta.Name("route.key.remove")), keyHandler.RemoveAPIKey)
	}

	dictGroup := api.Group("", ginx.M(meta.Group("route.dict.group"), meta.Roles(role.AdminRole)))
	dictHandler := handler.Dict
	{
		dictGroup.GET("/dict/info", ginx.M(meta.Name("route.dict.info"), meta.Roles(role.UserRole)), dictHandler.GetDictInfo)

		dictGroup.GET("/dict/list", ginx.M(meta.Name("route.dict.list")), dictHandler.ListDict)
		dictGroup.POST("/dict/create", ginx.M(meta.Name("route.dict.create")), dictHandler.CreateDict)
		dictGroup.POST("/dict/update", ginx.M(meta.Name("route.dict.update")), dictHandler.UpdateDict)
		dictGroup.DELETE("/dict/remove", ginx.M(meta.Name("route.dict.delete")), dictHandler.RemoveDict)

		dictGroup.GET("/dict/data/list", ginx.M(meta.Name("route.dict.data.list")), dictHandler.ListDictData)
		dictGroup.POST("/dict/data/create", ginx.M(meta.Name("route.dict.data.create")), dictHandler.CreateDictData)
		dictGroup.POST("/dict/data/update", ginx.M(meta.Name("route.dict.data.update")), dictHandler.UpdateDictData)
		dictGroup.DELETE("/dict/data/remove", ginx.M(meta.Name("route.dict.data.delete")), dictHandler.RemoveDictData)
	}
	return types.NopObj
}

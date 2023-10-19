package role

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

var RoleRouterSet = wire.NewSet(
	RoleProvider,
	wire.Struct(new(Handler), "*"),
	SetupRouter,
)

// HandlerRouter just for wire injection, no real influence
type HandlerRouter types.NopType

type Handler struct {
	Role RoleHandler
	Perm PermHandler
}

func SetupRouter(api *route.Router, handler Handler) HandlerRouter {
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
	}
	permGroup := api.Group("perm", nil)
	{
		permHandler := handler.Perm
		// GET
		permGroup.GET("list", nil, permHandler.GetPermList)
		// POST
		permGroup.POST("create", nil, permHandler.CreatePermission)
		permGroup.POST("update", nil, permHandler.UpdatePermission)
		// DELETE
		permGroup.DELETE("remove", nil, permHandler.RemovePermission)
	}
	return types.NopObj
}

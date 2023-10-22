package user

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/meta"
	"github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/pkg/ginx"
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
	Info   InfoHandler
	Modify ModifyHandler
	Roles  RoleHandler
}

func SetupRouter(api *ginx.RouterGroup, handler Handler) HandlerRouter {
	// user router
	userRouter := api.Group("user", ginx.M(meta.Group("route.user.group")))
	{
		infoHandler := handler.Info
		modifyHandler := handler.Modify
		roleHandler := handler.Roles

		// GET
		userRouter.GET("info", ginx.M(meta.Name("route.user.info"), meta.Roles(role.UserRole)), infoHandler.GetUserInfo)
		userRouter.GET("list", ginx.M(meta.Name("route.user.list"), meta.Roles(role.AdminRole)), infoHandler.GetUserInfoList)
		userRouter.GET("roles", ginx.M(meta.Name("route.user.getRoles"), meta.Roles(role.AdminRole)), roleHandler.GetUserRoles)

		// POST
		userRouter.POST("update", ginx.M(meta.Name("route.user.update"), meta.Roles(role.UserRole)), modifyHandler.UpdateUserInfo)
		userRouter.POST("create", ginx.M(meta.Name("route.user.create"), meta.Roles(role.AdminRole)), modifyHandler.CreateUser)
		userRouter.POST("roles", ginx.M(meta.Name("route.user.saveRoles"), meta.Roles(role.AdminRole)), roleHandler.SaveUserRoles)

		// DELETE
		userRouter.DELETE("remove", ginx.M(meta.Name("route.user.remove"), meta.Roles(role.AdminRole)), modifyHandler.RemoveUser)

	}

	return types.NopObj
}

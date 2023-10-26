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
	Info  InfoHandler
	Admin AdminHandler
}

func SetupRouter(api *ginx.RouterGroup, handler Handler) HandlerRouter {
	// user router
	userRouter := api.Group("/user", nil)
	{
		// user info handler
		infoHandler := handler.Info
		profileGroup := userRouter.Group("", ginx.M(meta.Group("route.user.profile.group"), meta.Roles(role.UserRole)))
		{
			profileGroup.GET("/profile", ginx.M(meta.Name("route.user.profile.info")), infoHandler.GetUserInfo)
			profileGroup.POST("/profile", ginx.M(meta.Name("route.user.profile.save")), infoHandler.UpdateUserInfo)
		}

		// user admin handler
		adminHandler := handler.Admin
		adminGroup := userRouter.Group("/admin", ginx.M(meta.Group("route.user.admin.group"), meta.Roles(role.AdminRole)))
		{
			adminGroup.GET("/list", ginx.M(meta.Name("route.user.admin.list")), adminHandler.GetUserInfoList)
			adminGroup.GET("/profile", ginx.M(meta.Name("route.user.admin.info")), adminHandler.GetSpecUserInfo)

			adminGroup.POST("/create", ginx.M(meta.Name("route.user.admin.create")), adminHandler.CreateUser)
			adminGroup.POST("/profile", ginx.M(meta.Name("route.user.admin.update")), adminHandler.SaveUser)

			adminGroup.DELETE("/remove", ginx.M(meta.Name("route.user.admin.remove")), adminHandler.RemoveUser)
		}
	}

	return types.NopObj
}

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
}

func SetupRouter(api *ginx.RouterGroup, handler Handler) HandlerRouter {
	// user router
	userRouter := api.Group("user", nil)
	{
		// user info
		infoRouter := userRouter.Group("", ginx.M(meta.Name("route.user")))
		infoHandler := handler.Info
		{
			infoRouter.GET("info", ginx.M(meta.Name("route.user.info"), meta.Anonymous), infoHandler.GetUserInfo)
			infoRouter.GET("list", ginx.M(meta.Name("route.user.list"), meta.Roles(role.AdminRole.Code)), infoHandler.GetUserInfoList)
		}

		// user modify
		modifyRouter := userRouter.Group("", ginx.M(meta.Name("route.user")))
		modifyHandler := handler.Modify
		{
			modifyRouter.POST("update", ginx.M(meta.Name("route.user.update")), modifyHandler.UpdateUserInfo)
			modifyRouter.DELETE("remove", ginx.M(meta.Name("route.user.remove"), meta.Roles(role.AdminRole.Code)), modifyHandler.RemoveUser)
		}

	}

	return types.NopObj
}

package user

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/route"
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

func SetupRouter(api *route.Router, handler Handler) HandlerRouter {
	// user router
	userRouter := api.Group("user", nil)
	{
		// user info
		infoRouter := userRouter.Group("", nil)
		infoHandler := handler.Info
		{
			infoRouter.GET("/user/info", route.Metas(), infoHandler.GetUserInfo)
			infoRouter.GET("/user/list", route.Metas(), infoHandler.GetUserInfoList)
		}

		// user modify
		modifyRouter := userRouter.Group("", nil)
		modifyHandler := handler.Modify
		{
			modifyRouter.POST("/user/update", route.Metas(), modifyHandler.UpdateUserInfo)
			modifyRouter.DELETE("/user/remove", route.Metas(), modifyHandler.RemoveUser)
		}

	}

	return types.NopObj
}

package system

import (
	"github.com/dstgo/wilson/app/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

var SystemApiSet = wire.NewSet(
	NewPingApi,
	NewAuthApi,
	NewEmailApi,
	NewRoleApi,
	NewSystemRouter,
)

type Router struct {
	PingApi  PingApi
	AuthApi  AuthApi
	EmailApi EmailApi
	RoleApi  RoleApi
}

func NewSystemRouter(root *route.Router, pingApi PingApi, authApi AuthApi, emailApi EmailApi, roleApi RoleApi) Router {

	// ping api
	{
		// GET
		root.GET("/ping", route.MetaSum(meta.NoAuth), pingApi.Ping)
	}

	// auth api
	{
		authGroup := root.Group("auth", nil)
		// POST
		authGroup.POST("/login", route.MetaSum(meta.NoAuth), authApi.Login)
		authGroup.POST("/register", route.MetaSum(meta.NoAuth), authApi.Register)

		// DELETE
		authGroup.DELETE("/logout", nil, authApi.Logout)
	}

	// email api
	{
		emailGroup := root.Group("email", nil)

		// GET
		emailGroup.GET("/code", route.MetaSum(meta.NoAuth), emailApi.SendCodeEmail)
	}

	return Router{PingApi: pingApi}
}

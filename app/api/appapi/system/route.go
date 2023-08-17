package system

import (
	"github.com/dstgo/wilson/app/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

var SystemApiSet = wire.NewSet(
	NewPingApi,
	NewAuthApi,
	NewSystemRouter,
)

type Router struct {
	PingApi PingApi
	AuthApi AuthApi
}

func NewSystemRouter(root *route.Router, pingApi PingApi, authApi AuthApi) Router {
	// ping api
	root.GET("/ping", route.Metas(meta.NoAuth), pingApi.Ping)

	// auth api
	authGroup := root.Group("auth", nil)
	authGroup.POST("/login", route.Metas(meta.NoAuth), authApi.Login)
	authGroup.POST("/register", route.Metas(meta.NoAuth), authApi.Register)
	authGroup.DELETE("/logout", nil, authApi.Logout)

	return Router{PingApi: pingApi}
}

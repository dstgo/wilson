package userApi

import (
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

var UserApiSet = wire.NewSet(
	NewUserInfoApi,
	NewUserRouter,
)

type Router struct {
	UserInfoApi *UserInfoApi
}

func NewUserRouter(root *route.Router, userInfo *UserInfoApi) *Router {

	return &Router{UserInfoApi: userInfo}
}

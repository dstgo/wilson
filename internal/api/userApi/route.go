package userApi

import (
	"github.com/dstgo/wilson/pkg/coco"
	"github.com/google/wire"
)

var UserApiSet = wire.NewSet(
	NewUserInfoApi,
	NewUserRouter,
)

type Router struct {
	UserInfoApi *UserInfoApi
}

func NewUserRouter(coco *coco.Core, userInfo *UserInfoApi) *Router {

	return &Router{UserInfoApi: userInfo}
}

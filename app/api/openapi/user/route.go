package user

import (
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

var UserApiSet = wire.NewSet(
	NewUserInfoApi,
	NewUserRouter,
)

func NewUserRouter(root *route.Router, info UserInfoApi) UserRouter {

	return UserRouter{info}
}

type UserRouter struct {
	UserInfoApi UserInfoApi
}

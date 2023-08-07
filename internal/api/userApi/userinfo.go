package userApi

import (
	"github.com/dstgo/wilson/internal/logic/userLogic"
)

func NewUserInfoApi(userL *userLogic.UserInfoLogic) *UserInfoApi {
	return &UserInfoApi{userLogic: userL}
}

type UserInfoApi struct {
	userLogic *userLogic.UserInfoLogic
}

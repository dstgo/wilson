package user

import "github.com/dstgo/wilson/app/logic/userLogic"

func NewUserInfoApi(userL userLogic.UserInfoLogic) UserInfoApi {
	return UserInfoApi{info: userL}
}

type UserInfoApi struct {
	info userLogic.UserInfoLogic
}

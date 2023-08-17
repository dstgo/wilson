package user

import (
	"github.com/dstgo/wilson/app/logic/userLogic"
	"github.com/gin-gonic/gin"
)

func NewUserInfoApi(userL userLogic.UserInfoLogic) UserInfoApi {
	return UserInfoApi{info: userL}
}

type UserInfoApi struct {
	info userLogic.UserInfoLogic
}

func (ui UserInfoApi) GetUserInfo(ctx *gin.Context) {

}

func (ui UserInfoApi) GetUserInfoById(ctx *gin.Context) {

}

func (ui UserInfoApi) GetUserInfoList(ctx *gin.Context) {

}

func (ui UserInfoApi) UpdateUserInfo(ctx *gin.Context) {

}

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserProviderSet = wire.NewSet(
	NewInfoData,
	NewInfoLogic,
	NewInfoHandler,
)

func NewInfoHandler(userL InfoLogic) InfoHandler {
	return InfoHandler{info: userL}
}

type InfoHandler struct {
	info InfoLogic
}

func (ui InfoHandler) GetUserInfo(ctx *gin.Context) {

}

func (ui InfoHandler) GetUserInfoById(ctx *gin.Context) {

}

func (ui InfoHandler) GetUserInfoList(ctx *gin.Context) {

}

func (ui InfoHandler) UpdateUserInfo(ctx *gin.Context) {

}

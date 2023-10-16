package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserProviderSet = wire.NewSet(
	NewInfoData,
	NewUserInfo,
	NewInfoHandler,
)

func NewInfoHandler(userL UserInfo) InfoHandler {
	return InfoHandler{info: userL}
}

type InfoHandler struct {
	info UserInfo
}

// GetUserInfo
//
//	@Summary		GetUserInfo
//	@Description	get specific user simple info
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/user/info [GET]
func (ui InfoHandler) GetUserInfo(ctx *gin.Context) {

}

// UpdateUserInfo
//
//	@Summary		UpdateUserInfo
//	@Description	update the specific user info
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/user/update [GET]
func (ui InfoHandler) UpdateUserInfo(ctx *gin.Context) {

}

// RemoveUser
//
//	@Summary		RemoveUser
//	@Description	Remove the specific user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/user/remove [get]
func (ui InfoHandler) RemoveUser(ctx *gin.Context) {

}

// GetUserInfoList
//
//	@Summary		GetUserInfoList
//	@Description	get specific user list
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/user/list [GET]
func (ui InfoHandler) GetUserInfoList(ctx *gin.Context) {

}

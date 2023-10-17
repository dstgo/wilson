package user

import (
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserApiProviderSet = wire.NewSet(
	NewInfoLogic,
	NewInfoApi,
	user.NewUserData,
	user.NewUserInfo,
)

func NewInfoApi(info InfoLogic, hInfo user.UserInfo) InfoApi {
	return InfoApi{
		info:  info,
		hInfo: hInfo,
	}
}

type InfoApi struct {
	info  InfoLogic
	hInfo user.UserInfo
}

// UserInfo
//
//	@Summary		GetUserInfoById
//	@Description	get user simple info
//	@Tags           user
//	@Accept			query
//	@Produce		json
//	@Param			name	query		string	true	"comment"
//	@Router			/user/info [get]
func (i InfoApi) UserInfo(ctx *gin.Context) {

}

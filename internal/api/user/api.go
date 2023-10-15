package user

import (
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserApiProviderSet = wire.NewSet(
	NewInfoLogic,
	NewInfoApi,
	user.NewInfoData,
	user.NewInfoLogic,
)

func NewInfoApi(info InfoLogic, hInfo user.InfoLogic) InfoApi {
	return InfoApi{
		info:  info,
		hInfo: hInfo,
	}
}

type InfoApi struct {
	info  InfoLogic
	hInfo user.InfoLogic
}

// UserInfo
//
//	@Summary		path params example
//	@Description	get user simple info
//	@Tags
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	true	"comment"
//	@Router			/open/user/info [get]
func (i InfoApi) UserInfo(ctx *gin.Context) {

}

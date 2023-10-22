package user

import (
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserApiProviderSet = wire.NewSet(
	NewInfoLogic,
	NewInfoApi,
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
// @Summary      UserInfo
// @Description  get user info
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uuid query     types.Uid  true  "user uuid"
// @Success      200  {object}  types.Response{data=user.Info}
// @Router       /api [GET]
func (i InfoApi) UserInfo(ctx *gin.Context) {

}

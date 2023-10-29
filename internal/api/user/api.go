package user

import (
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/dstgo/wilson/internal/types/auth"
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

// KeyInfo
// @Summary      KeyInfo
// @Description  get the key info
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response{data=auth.APIKey}
// @Router       /info [GET]
// @Security    ApiKeyAuth
func (i InfoApi) KeyInfo(ctx *gin.Context) {
	info := authen.GetContextKeyInfo(ctx)
	resp.Ok(ctx).MsgI18n("op.query.ok").Data(auth.APIKey{
		Key:       info.Key,
		Name:      info.Name,
		ExpiredAt: info.ExpiredAt,
	}).Send()
}

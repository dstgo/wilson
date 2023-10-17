package user

import (
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/core/valid"
	"github.com/dstgo/wilson/internal/types/api"
	"github.com/dstgo/wilson/internal/types/api/user"
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
// @Summary      GetUserInfo
// @Description  get specific user simple info
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id      query   int     true    "userId"
// @Success      200  {object}  api.Response{data=user.Info}
// @Router       /user/info [GET]
func (ui InfoHandler) GetUserInfo(ctx *gin.Context) {
	var id api.Id
	if err := valid.BindAndResp(ctx, valid.Query(&id)); err != nil {
		return
	}

	info, err := ui.info.GetUserInfo(id.Uint())
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("user.findfail").Send()
		return
	}
	resp.Ok(ctx).Data(info).MsgI18n("user.findok").Send()
}

// GetUserInfoList
// @Summary      GetUserInfoList
// @Description  get specific user list
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userPageOptIon	body	user.PageOption	true	"comment"
// @Success      200  {object}  api.Response{data=[]user.Info}
// @Router       /user/list [GET]
func (ui InfoHandler) GetUserInfoList(ctx *gin.Context) {
	var page user.PageOption
	if err := valid.Bind(ctx, valid.Query(&page)); err != nil {
		return
	}

	list, err := ui.info.GetUserInfoList(page)
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("user.findFail").Send()
		return
	}
	resp.Ok(ctx).Data(list).MsgI18n("user.findOK").Send()
}

// UpdateUserInfo
// @Summary      UpdateUserInfo
// @Description  update the specific user info
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        updateInfoOption	body	user.UpdateInfoOption	true	"comment"
// @Success      200  {object}  api.Response
// @Router       /user/update [POST]
func (ui InfoHandler) UpdateUserInfo(ctx *gin.Context) {
	var updateUserOpt user.UpdateInfoOption
	if err := valid.BindAndResp(ctx, valid.Json(&updateUserOpt)); err != nil {
		return
	}

	if err := ui.info.UpdateUserInfo(updateUserOpt); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("user.updateFail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("user.updateOk").Send()
}

// RemoveUser
// @Summary      RemoveUser
// @Description  Remove the specific user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  query   int true    "userId"
// @Success      200  {object}  api.Response
// @Router       /user/remove [DELETE]
func (ui InfoHandler) RemoveUser(ctx *gin.Context) {
	var id api.Id
	if err := valid.BindAndResp(ctx, valid.Query(&id)); err != nil {
		return
	}
	if err := ui.info.RemoveUser(id.Uint()); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("user.removeFail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("user.removeOk").Send()
}

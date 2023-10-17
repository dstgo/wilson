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
	NewUserData,
	NewUserInfo,
	NewUserModify,
	NewInfoHandler,
	NewModifyHandler,
)

func NewInfoHandler(info UserInfo) InfoHandler {
	return InfoHandler{info: info}
}

type InfoHandler struct {
	info UserInfo
}

// GetUserInfo
// @Summary      GetUserInfo
// @Description  get specific user simple info
// @Tags         user/info
// @Accept       json
// @Produce      json
// @Param        uuid      query   api.UUID     true    "user unique id"
// @Success      200  {object}  api.Response{data=user.Info}
// @Router       /user/info [GET]
// @security BearerAuth
func (ui InfoHandler) GetUserInfo(ctx *gin.Context) {
	var uuid api.UUID
	if err := valid.BindAndResp(ctx, valid.Query(&uuid)); err != nil {
		return
	}

	info, err := ui.info.GetUserInfoByUUID(uuid.UUID)
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("user.findfail").Send()
		return
	}
	resp.Ok(ctx).Data(info).MsgI18n("user.findok").Send()
}

// GetUserInfoList
// @Summary      GetUserInfoList
// @Description  get specific user list
// @Tags         user/info
// @Accept       json
// @Produce      json
// @Param        userPageOptIon	query	user.PageOption	true	"comment"
// @Success      200  {object}  api.Response{data=[]user.Info}
// @Router       /user/list [GET]
// @security BearerAuth
func (ui InfoHandler) GetUserInfoList(ctx *gin.Context) {
	var page user.PageOption
	if err := valid.BindAndResp(ctx, valid.Query(&page)); err != nil {
		return
	}

	list, err := ui.info.GetUserInfoList(page)
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("user.findFail").Send()
		return
	}
	resp.Ok(ctx).Data(list).MsgI18n("user.findOK").Send()
}

func NewModifyHandler(modify UserModify) ModifyHandler {
	return ModifyHandler{modify: modify}
}

type ModifyHandler struct {
	modify UserModify
}

// UpdateUserInfo
// @Summary      UpdateUserInfo
// @Description  update the specific user info
// @Tags         user/modify
// @Accept       json
// @Produce      json
// @Param        uuid   query      api.UUID  true  "uuid"
// @Param        updateInfoOption	body	user.UpdateInfoOption	true	"comment"
// @Success      200  {object}  api.Response
// @Router       /user/update [POST]
// @security BearerAuth
func (ui ModifyHandler) UpdateUserInfo(ctx *gin.Context) {
	var (
		updateUserOpt user.UpdateInfoOption
		uuid          api.UUID
	)

	if err := valid.BindAndResp(ctx,
		valid.Query(&uuid),
		valid.Json(&updateUserOpt)); err != nil {
		return
	}

	updateUserOpt.UUID = uuid.UUID

	if err := ui.modify.Update(updateUserOpt); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("user.updateFail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("user.updateOk").Send()
}

// RemoveUser
// @Summary      RemoveUser
// @Description  Remove the specific user
// @Tags         user/modify
// @Accept       json
// @Produce      json
// @Param        uuid  query   api.UUID true    "uuid"
// @Success      200  {object}  api.Response
// @Router       /user/remove [DELETE]
// @security BearerAuth
func (ui ModifyHandler) RemoveUser(ctx *gin.Context) {
	var uuid api.UUID
	if err := valid.BindAndResp(ctx, valid.Query(&uuid)); err != nil {
		return
	}
	if err := ui.modify.Remove(uuid.UUID); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("user.removeFail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("user.removeOk").Send()
}

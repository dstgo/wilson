package user

import (
	"github.com/dstgo/wilson/internal/core/bind"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/user"
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
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uuid      query    types.Uid     true    "user unique id"
// @Success      200  {object}  types.Response{data=user.Info}
// @Router       /user/info [GET]
// @security BearerAuth
func (ui InfoHandler) GetUserInfo(ctx *gin.Context) {
	var uuid types.Uid
	if err := bind.BindAndResp(ctx, bind.Query(&uuid)); err != nil {
		return
	}

	info, err := ui.info.GetUserInfoByUUID(uuid.UUID)
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.query.fail").Send()
		return
	}
	resp.Ok(ctx).Data(info).MsgI18n("op.query.ok").Send()
}

// GetUserInfoList
// @Summary      GetUserInfoList
// @Description  get specific user list
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userPageOptIon	query	user.PageOption	true	"comment"
// @Success      200  {object}  types.Response{data=[]user.Info}
// @Router       /user/list [GET]
// @security BearerAuth
func (ui InfoHandler) GetUserInfoList(ctx *gin.Context) {
	var page user.PageOption
	if err := bind.BindAndResp(ctx, bind.Query(&page)); err != nil {
		return
	}

	list, err := ui.info.GetUserInfoList(page)
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.query.fail").Send()
		return
	}
	resp.Ok(ctx).Data(list).MsgI18n("op.query.ok").Send()
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
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uuid   query      types.Uid  true  "uuid"
// @Param        updateInfoOption	body	user.UpdateInfoOption	true	"comment"
// @Success      200  {object}  types.Response
// @Router       /user/update [POST]
// @security BearerAuth
func (ui ModifyHandler) UpdateUserInfo(ctx *gin.Context) {
	var (
		updateUserOpt user.UpdateInfoOption
		uuid          types.Uid
	)

	if err := bind.BindAndResp(ctx,
		bind.Query(&uuid),
		bind.Json(&updateUserOpt)); err != nil {
		return
	}

	updateUserOpt.UUID = uuid.UUID

	if err := ui.modify.Update(updateUserOpt); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.update.fail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.update.fail").Send()
}

// RemoveUser
// @Summary      RemoveUser
// @Description  Remove the specific user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uuid  query   types.Uid  true    "uuid"
// @Success      200  {object}  types.Response
// @Router       /user/remove [DELETE]
// @security BearerAuth
func (ui ModifyHandler) RemoveUser(ctx *gin.Context) {
	var uuid types.Uid
	if err := bind.BindAndResp(ctx, bind.Query(&uuid)); err != nil {
		return
	}
	if err := ui.modify.Remove(uuid.UUID); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.delete.fail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.delete.ok").Send()
}

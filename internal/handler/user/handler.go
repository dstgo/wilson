package user

import (
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/user"
	"github.com/dstgo/wilson/pkg/ginx/bind"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserProviderSet = wire.NewSet(
	NewUserInfo,
	NewUserModify,
	NewInfoHandler,
	NewAdminHandler,
)

func NewInfoHandler(info UserInfo, modify UserModify) InfoHandler {
	return InfoHandler{info: info, modify: modify}
}

type InfoHandler struct {
	info   UserInfo
	modify UserModify
}

// GetUserInfo
// @Summary      GetUserInfo
// @Description  [user]
// @Description  get own user info
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response{data=user.Info}
// @Router       /user/profile [GET]
// @security BearerAuth
func (i InfoHandler) GetUserInfo(ctx *gin.Context) {

	contextTokenInfo := authen.GetContextTokenInfo(ctx)

	info, err := i.info.GetUserInfoByUUID(contextTokenInfo.UUID)
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.query.fail").Send()
		return
	}
	resp.Ok(ctx).Data(info).MsgI18n("op.query.ok").Send()
}

// UpdateUserInfo
// @Summary      UpdateUserInfo
// @Description  [user]
// @Description  update own user info
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        updateInfoOption	body	user.UpdateInfoOption	true	"updateInfoOption"
// @Success      200  {object}  types.Response
// @Router       /user/profile [POST]
// @security BearerAuth
func (i InfoHandler) UpdateUserInfo(ctx *gin.Context) {
	var (
		updateUserOpt user.UpdateInfoOption
	)

	if err := bind.Binds(ctx,
		bind.Json(&updateUserOpt)); err != nil {
		return
	}

	// get user info from context
	info := authen.GetContextTokenInfo(ctx)
	updateUserOpt.UUID = info.UUID

	if err := i.modify.Update(updateUserOpt); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.update.fail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.update.ok").Send()
}

func NewAdminHandler(info UserInfo, modify UserModify) AdminHandler {
	return AdminHandler{info: info, modify: modify}
}

type AdminHandler struct {
	info   UserInfo
	modify UserModify
}

// GetSpecUserInfo
// @Summary      GetSpecUserInfo
// @Description  [admin]
// @Description  get specified user information
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uuid query     types.Uid  true  "user uuid"
// @Success      200  {object}  types.Response
// @Router       /user/admin/profile [GET]
// @security BearerAuth
func (a AdminHandler) GetSpecUserInfo(ctx *gin.Context) {
	var uuid types.Uid
	if err := bind.Binds(ctx, bind.Query(&uuid)); err != nil {
		return
	}

	info, err := a.info.GetUserInfoByUUID(uuid.UUID)
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.query.fail").Send()
		return
	}
	resp.Ok(ctx).Data(info).MsgI18n("op.query.ok").Send()
}

// GetUserInfoList
// @Summary      GetUserInfoList
// @Description  [admin]
// @Description  get specific user list
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userPageOptIon	query	user.PageOption	true	"comment"
// @Success      200  {object}  types.Response{data=[]user.Info}
// @Router       /user/admin/list [GET]
// @security BearerAuth
func (a AdminHandler) GetUserInfoList(ctx *gin.Context) {
	var page user.PageOption
	if err := bind.Binds(ctx, bind.Query(&page)); err != nil {
		return
	}

	list, err := a.info.GetUserInfoList(page)
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.query.fail").Send()
		return
	}
	resp.Ok(ctx).Data(list).MsgI18n("op.query.ok").Send()
}

// CreateUser
// @Summary      CreateUser
// @Description  [admin]
// @Description  create new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        createOpt      body     user.CreateUserOption  true  "CreateUserOption"
// @Success      200  {object}  types.Response
// @Router       /user/admin/create [POST]
// @security BearerAuth
func (a AdminHandler) CreateUser(ctx *gin.Context) {
	var createOpt user.CreateUserOption
	if err := bind.Binds(ctx, bind.Json(&createOpt)); err != nil {
		return
	}

	if err := a.modify.Create(createOpt); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.create.fail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.create.ok").Send()
}

// SaveUser
// @Summary      SaveUser
// @Description  [admin]
// @Description  save specified user information
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        updateInfoOption	body	user.SaveUserDetailOption	true	"updateInfoOption"
// @Success      200  {object}  types.Response
// @Router       /user/admin/profile [POST]
// @security BearerAuth
func (a AdminHandler) SaveUser(ctx *gin.Context) {
	var saveOption user.SaveUserDetailOption
	if err := bind.Binds(ctx, bind.Json(&saveOption)); err != nil {
		return
	}

	err := a.modify.Save(saveOption)
	if err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.update.fail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.update.ok").Send()
}

// RemoveUser
// @Summary      RemoveUser
// @Description  [admin]
// @Description  Remove the specific user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uuid  query   types.Uid  true    "uuid"
// @Success      200  {object}  types.Response
// @Router       /user/admin/remove [DELETE]
// @security BearerAuth
func (a AdminHandler) RemoveUser(ctx *gin.Context) {
	var uuid types.Uid
	if err := bind.Binds(ctx, bind.Query(&uuid)); err != nil {
		return
	}
	if err := a.modify.Remove(uuid.UUID); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.delete.fail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.delete.ok").Send()
}

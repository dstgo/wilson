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
	NewUserInfo,
	NewUserModify,
	NewInfoHandler,
	NewModifyHandler,
	NewUserRole,
	NewUserRoleHandler,
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
	if err := bind.Binds(ctx, bind.Query(&uuid)); err != nil {
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
	if err := bind.Binds(ctx, bind.Query(&page)); err != nil {
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
// @Param        updateInfoOption	body	user.UpdateInfoOption	true	"comment"
// @Success      200  {object}  types.Response
// @Router       /user/update [POST]
// @security BearerAuth
func (ui ModifyHandler) UpdateUserInfo(ctx *gin.Context) {
	var (
		updateUserOpt user.UpdateInfoOption
	)

	if err := bind.Binds(ctx,
		bind.Json(&updateUserOpt)); err != nil {
		return
	}

	if err := ui.modify.Update(updateUserOpt); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.update.fail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.update.fail").Send()
}

// CreateUser
// @Summary      CreateUser
// @Description  create new user
// @Tags         app
// @Accept       json
// @Produce      json
// @Param        createOpt      body     user.CreateUserOption  true  "CreateUserOption"
// @Success      200  {object}  types.Response
// @Router       /api [GET]
func (ui ModifyHandler) CreateUser(ctx *gin.Context) {
	var createOpt user.CreateUserOption
	if err := bind.Binds(ctx, bind.Json(&createOpt)); err != nil {
		return
	}

	if err := ui.modify.Create(createOpt); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.create.fail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.create.fail").Send()
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
	if err := bind.Binds(ctx, bind.Query(&uuid)); err != nil {
		return
	}
	if err := ui.modify.Remove(uuid.UUID); err != nil {
		resp.Fail(ctx).Error(err).MsgI18n("op.delete.fail").Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.delete.ok").Send()
}

func NewUserRoleHandler(role UserRole) RoleHandler {
	return RoleHandler{role: role}
}

type RoleHandler struct {
	role UserRole
}

// GetUserRoles
// @Summary      GetUserRoles
// @Description  get user roles
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uuid query     types.Uid  true  "user uuid"
// @Success      200  {object}  types.Response{data=[]role.RoleInfo}
// @Router       /user/roles [GET]
func (u RoleHandler) GetUserRoles(ctx *gin.Context) {
	var uuid types.Uid
	if err := bind.Binds(ctx, bind.Query(&uuid)); err != nil {
		return
	}
	roles, err := u.role.GetUserRoles(uuid.UUID)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.query.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.query.ok").Data(roles).Send()
}

// SaveUserRoles
// @Summary      SaveUserRoles
// @Description  get string by ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        saveOption     body      user.SaveRoleOption  true  "SaveRoleOption"
// @Success      200  {object}  types.Response
// @Router       /user/roles [POST]
func (u RoleHandler) SaveUserRoles(ctx *gin.Context) {
	var saveOpt user.SaveRoleOption
	if err := bind.Binds(ctx, bind.Json(&saveOpt)); err != nil {
		return
	}

	err := u.role.SaveRoles(saveOpt.UUID, saveOpt.RoleIds)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.update.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.update.ok").Send()
}

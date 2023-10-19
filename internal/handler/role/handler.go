package role

import (
	"github.com/dstgo/wilson/internal/core/bind"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/types/api"
	"github.com/dstgo/wilson/internal/types/api/role"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RoleProvider = wire.NewSet(
	NewPermData,
	NewRoleData,
	NewPermService,
	NewRoleService,
	NewRoleHandler,
	NewPermHandler,
)

func NewRoleHandler(role RoleService) RoleHandler {
	return RoleHandler{role: role}
}

type RoleHandler struct {
	role RoleService
}

// GetRoleList
// @Summary      GetRoleList
// @Description  get role list by page
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        page  query    role.PageOption  true  "page option"
// @Success      200  {object}  api.Response{data=[]role.RoleInfo}
// @Router       /role/list [GET]
func (r RoleHandler) GetRoleList(ctx *gin.Context) {
	var pageOpt role.PageOption
	if err := bind.BindAndResp(ctx, bind.Query(&pageOpt)); err != nil {
		return
	}

	roleInfos, err := r.role.ListRoleByPage(pageOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.query.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.query.ok").Data(roleInfos).Send()
}

// GetRolePerms
// @Summary      GetRolePerms
// @Description  get permissions list belong to role
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        queryOpt   query   api.Id true  "role perms query opt"
// @Success      200  {object}  api.Response{data=[]role.PermGroup}
// @Router       /role/perms [GET]
func (r RoleHandler) GetRolePerms(ctx *gin.Context) {
	var queryOpt api.Id
	if err := bind.BindAndResp(ctx, bind.Query(&queryOpt)); err != nil {
		return
	}

	perms, err := r.role.GetRolePerms(queryOpt.Uint())
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.query.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.query.ok").Data(perms).Send()
}

// CreateRole
// @Summary      CreateRole
// @Description  create a new role
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        createRole   body      role.CreateRoleOption  true  "create role"
// @Success      200  {object}  api.Response
// @Router       /role/create [POST]
func (r RoleHandler) CreateRole(ctx *gin.Context) {
	var createOpt role.CreateRoleOption
	if err := bind.BindAndResp(ctx, bind.Json(&createOpt)); err != nil {
		return
	}
	err := r.role.CreateRole(createOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.create.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.create.ok").Send()
}

// UpdateRole
// @Summary      UpdateRole
// @Description  update the specified role info
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        updateRole   body      role.UpdateRoleOption  true  "update role"
// @Success      200  {object}  api.Response
// @Router       /role/update [POST]
func (r RoleHandler) UpdateRole(ctx *gin.Context) {
	var updateOpt role.UpdateRoleOption
	if err := bind.BindAndResp(ctx, bind.Json(&updateOpt)); err != nil {
		return
	}
	err := r.role.UpdateRole(updateOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.update.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.update.ok").Send()
}

// GrantRolePerms
// @Summary      GrantRolePerms
// @Description  grant permissions for the specified role
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        GrantOption   body    role.GrantOption  true  "grant role"
// @Success      200  {object}  api.Response
// @Router       /role/grant [POST]
func (r RoleHandler) GrantRolePerms(ctx *gin.Context) {
	var grantOption role.GrantOption
	if err := bind.BindAndResp(ctx, bind.Json(&grantOption)); err != nil {
		return
	}

	err := r.role.GrantRolePerms(grantOption.RoleId, grantOption.Tag, grantOption.PermId)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.update.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.update.ok").Send()
}

// RemoveRole
// @Summary      RemoveRole
// @Description  remove a role,and its permission record will be deleted too
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        id   query     api.Id  true  "roleD id"
// @Success      200  {object}  api.Response
// @Router       /role/remove [DELETE]
func (r RoleHandler) RemoveRole(ctx *gin.Context) {
	var roleId api.Id
	if err := bind.BindAndResp(ctx, bind.Query(&roleId)); err != nil {
		return
	}

	err := r.role.RemoveRole(roleId.Uint())
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.delete.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.delete.ok").Send()
}

func NewPermHandler(perm PermService) PermHandler {
	return PermHandler{perm: perm}
}

type PermHandler struct {
	perm PermService
}

// GetPermList
// @Summary      GetPermList
// @Description  Get Permission list by page
// @Tags         perm
// @Accept       json
// @Produce      json
// @Param        page  query    role.PageOption  true  "page option"
// @Success      200  {object}  api.Response{data=[]role.PermInfo}
// @Router       /perm/list [GET]
func (r PermHandler) GetPermList(ctx *gin.Context) {
	var pageOpt role.PageOption
	if err := bind.BindAndResp(ctx, bind.Query(&pageOpt)); err != nil {
		return
	}

	roleInfos, err := r.perm.GetPermList(pageOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.query.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.query.ok").Data(roleInfos).Send()
}

// CreatePermission
// @Summary      CreatePermission
// @Description  create a new permission
// @Tags         perm
// @Accept       json
// @Produce      json
// @Param        createPerm   body      role.CreatePermOption  true  "create perm"
// @Success      200  {object}  api.Response
// @Router       /perm/create [POST]
func (r PermHandler) CreatePermission(ctx *gin.Context) {
	var createOpt role.CreatePermOption
	if err := bind.BindAndResp(ctx, bind.Json(&createOpt)); err != nil {
		return
	}

	err := r.perm.CreatePerm(createOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.create.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.create.ok").Send()
}

// UpdatePermission
// @Summary      UpdatePermission
// @Description  update the specified permission info
// @Tags         perm
// @Accept       json
// @Produce      json
// @Param        updatePermOpt   body      role.UpdatePermOption  true  "update perm"
// @Success      200  {object}  api.Response
// @Router       /perm/update [POST]
func (r PermHandler) UpdatePermission(ctx *gin.Context) {
	var updateOpt role.UpdatePermOption
	if err := bind.BindAndResp(ctx, bind.Json(&updateOpt)); err != nil {
		return
	}
	err := r.perm.UpdatePerm(updateOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.update.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.update.ok").Send()
}

// RemovePermission
// @Summary      RemovePermission
// @Description  remove the specified permission
// @Tags         perm
// @Accept       json
// @Produce      json
// @Param        id   query     api.Id  true  "perm id"
// @Success      200  {object}  api.Response
// @Router       /perm/remove [DELETE]
func (r PermHandler) RemovePermission(ctx *gin.Context) {
	var permId api.Id
	if err := bind.BindAndResp(ctx, bind.Query(&permId)); err != nil {
		return
	}

	err := r.perm.RemovePerm(permId.Uint())
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.delete.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.delete.ok").Send()
}

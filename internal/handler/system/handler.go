package system

import (
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/core/bind"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/data/cache"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/auth"
	"github.com/dstgo/wilson/internal/types/code"
	"github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SystemProviderSet = wire.NewSet(
	cache.TokenCacheProvider,
	NewPingHandler,
	NewPingLogic,
	NewAuthenticator,
	NewAuthHandler,
	NewRoleEnforcer,
	NewRoleHandler,
)

func NewPingHandler(logic PingApp) PingHandler {
	return PingHandler{
		PingLogic: logic,
	}
}

type PingHandler struct {
	PingLogic PingApp
}

// Ping
// @Summary      Ping
// @Description  test app api if is accessible
// @Tags         system
// @Accept       json
// @Produce      json
// @Param        name	query	system.PingRequest	true	"ping name"
// @Success      200  {object}  types.Response{data=auth.PingReply}
// @Router       /ping [GET]
func (p PingHandler) Ping(ctx *gin.Context) {
	pingReq := new(system.PingRequest)
	err := bind.BindAndResp(ctx,
		bind.Query(pingReq),
	)
	if err != nil {
		return
	}

	reply := p.PingLogic.Ping(pingReq.Name)

	resp.Ok(ctx).Code(code.RequestOk).Msg("pong").Data(reply).Send()
}

// Pong
// @Summary      Pong
// @Description  test app api authentication if is work
// @Tags         system
// @Accept       json
// @Produce      json
// @Param        name   query      system.PingRequest  true  "pong name"
// @Success      200  {object}  types.Response
// @Router       /pong [GET]
func (p PingHandler) Pong(ctx *gin.Context) {
	pongReq := new(system.PingRequest)
	err := bind.BindAndResp(ctx,
		bind.Query(pongReq),
	)
	if err != nil {
		return
	}

	reply := p.PingLogic.Pong(pongReq.Name)

	resp.Ok(ctx).Code(code.RequestOk).Msg("ping").Data(reply).Send()
}

func NewAuthHandler(authen Authenticator) AuthHandler {
	return AuthHandler{Authlogic: authen}
}

type AuthHandler struct {
	Authlogic Authenticator
}

// Login
// @Summary      Login
// @Description  if login success, return jwt token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginBody	body	auth.LoginOption	true	"comment"
// @Success      200  {object}  types.Response{data=auth.Token}
// @Router       /auth/login [POST]
func (a AuthHandler) Login(ctx *gin.Context) {
	loginRequest := new(auth.LoginOption)
	if err := bind.BindAndResp(ctx, bind.Json(loginRequest)); err != nil {
		return
	}
	signedJwt, err := a.Authlogic.TryLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		resp.Fail(ctx).Code(code.LoginFailed).MsgI18n("auth.loginFail").Error(err).Send()
		return
	}
	resp.Ok(ctx).Code(code.LoginOk).MsgI18n("auth.loginOk").
		Data(auth.Token{Token: signedJwt.SignedJwt}).Send()
}

// Register
// @Summary      Register
// @Description  user register api
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        registerBody	body	auth.RegisterOption	true	"comment"
// @Success      200  {object}  types.Response
// @Router       /auth/register [POST]
func (a AuthHandler) Register(ctx *gin.Context) {
	registerRequest := new(auth.RegisterOption)
	if err := bind.BindAndResp(ctx, bind.Json(registerRequest)); err != nil {
		return
	}
	err := a.Authlogic.TryRegisterNewUser(registerRequest.Username, registerRequest.Password, registerRequest.Code)
	if err != nil {
		resp.Fail(ctx).Code(code.RegisterFailed).MsgI18n("auth.registerFail").Error(err).Send()
		return
	}
	resp.Ok(ctx).Code(code.RegisterOk).MsgI18n("auth.registerOk").Send()
}

// Logout
// @Summary      Logout
// @Description  user logout
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /auth/logout [DELETE]
// @security BearerAuth
func (a AuthHandler) Logout(ctx *gin.Context) {
	// get user info from parsed request context
	tokenInfo := authen.GetContextTokenInfo(ctx)
	err := a.Authlogic.TryLogout(tokenInfo.ID)
	if err != nil {
		resp.Fail(ctx).Code(code.LogoutFailed).MsgI18n("auth.logoutFail").Error(err).Send()
		return
	}
	resp.Ok(ctx).Code(code.LogoutOK).MsgI18n("auth.logoutOk").Send()
}

// ForgotPassword
// @Summary      ForgotPassword
// @Description  forgot password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        form   body      auth.ForgotPasswordOption  true  "Account ID"
// @Success      200  {object}  types.Response
// @Router       /auth/forgotpwd [POST]
func (a AuthHandler) ForgotPassword(ctx *gin.Context) {
	changePasswordReq := new(auth.ForgotPasswordOption)
	err := bind.BindAndResp(ctx,
		bind.Json(changePasswordReq),
	)
	if err != nil {
		return
	}

	err = a.Authlogic.ChangePassword(changePasswordReq.Password, changePasswordReq.Code)
	if err != nil {
		resp.Fail(ctx).MsgI18n("auth.changePasswdFail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("auth.changePasswdOk").Send()
}

func NewRoleHandler(enforcer RoleEnforcer) RoleHandler {
	return RoleHandler{enforcer: enforcer}
}

type RoleHandler struct {
	enforcer RoleEnforcer
}

// GetRoleList
// @Summary      GetRoleList
// @Description  get role list by page
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        page  query    role.PageOption  true  "page option"
// @Success      200  {object}  types.Response{data=[]role.RoleInfo}
// @Router       /role/list [GET]
func (r RoleHandler) GetRoleList(ctx *gin.Context) {
	var pageOpt role.PageOption
	if err := bind.BindAndResp(ctx, bind.Query(&pageOpt)); err != nil {
		return
	}

	roleInfos, err := r.enforcer.ListRole(pageOpt)
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
// @Param        queryOpt   query   types.Id true  "role perms query opt"
// @Success      200  {object}  types.Response{data=[]role.PermGroup}
// @Router       /role/perms [GET]
// @security BearerAuth
func (r RoleHandler) GetRolePerms(ctx *gin.Context) {
	var queryOpt types.Id
	if err := bind.BindAndResp(ctx, bind.Query(&queryOpt)); err != nil {
		return
	}

	perms, err := r.enforcer.ListRolePerms(queryOpt.Uint())
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
// @Success      200  {object}  types.Response
// @Router       /role/create [POST]
// @security BearerAuth
func (r RoleHandler) CreateRole(ctx *gin.Context) {
	var createOpt role.CreateRoleOption
	if err := bind.BindAndResp(ctx, bind.Json(&createOpt)); err != nil {
		return
	}
	err := r.enforcer.CreateRole(createOpt)
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
// @Success      200  {object}  types.Response
// @Router       /role/update [POST]
// @security BearerAuth
func (r RoleHandler) UpdateRole(ctx *gin.Context) {
	var updateOpt role.UpdateRoleOption
	if err := bind.BindAndResp(ctx, bind.Json(&updateOpt)); err != nil {
		return
	}
	err := r.enforcer.UpdateRole(updateOpt)
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
// @Success      200  {object}  types.Response
// @Router       /role/grant [POST]
// @security BearerAuth
func (r RoleHandler) GrantRolePerms(ctx *gin.Context) {
	var grantOption role.GrantOption
	if err := bind.BindAndResp(ctx, bind.Json(&grantOption)); err != nil {
		return
	}

	err := r.enforcer.UpdateRolePerms(grantOption)
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
// @Param        id   query     types.Id  true  "roleD id"
// @Success      200  {object}  types.Response
// @Router       /role/remove [DELETE]
// @security BearerAuth
func (r RoleHandler) RemoveRole(ctx *gin.Context) {
	var roleId types.Id
	if err := bind.BindAndResp(ctx, bind.Query(&roleId)); err != nil {
		return
	}

	err := r.enforcer.RemoveRole(roleId.Uint())
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.delete.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.delete.ok").Send()
}

// GetPermList
// @Summary      GetPermList
// @Description  Get Permission list by page
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        page  query    role.PageOption  true  "page option"
// @Success      200  {object}  types.Response{data=[]role.PermInfo}
// @Router       /perm/list [GET]
// @security BearerAuth
func (r RoleHandler) GetPermList(ctx *gin.Context) {
	var pageOpt role.PageOption
	if err := bind.BindAndResp(ctx, bind.Query(&pageOpt)); err != nil {
		return
	}

	roleInfos, err := r.enforcer.ListPerms(pageOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.query.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.query.ok").Data(roleInfos).Send()
}

// CreatePermission
// @Summary      CreatePermission
// @Description  create a new permission
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        createPerm   body      role.CreatePermOption  true  "create perm"
// @Success      200  {object}  types.Response
// @Router       /perm/create [POST]
// @security BearerAuth
func (r RoleHandler) CreatePermission(ctx *gin.Context) {
	var createOpt role.CreatePermOption
	if err := bind.BindAndResp(ctx, bind.Json(&createOpt)); err != nil {
		return
	}

	err := r.enforcer.CreatePerm(createOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.create.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.create.ok").Send()
}

// UpdatePermission
// @Summary      UpdatePermission
// @Description  update the specified permission info
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        updatePermOpt   body      role.UpdatePermOption  true  "update perm"
// @Success      200  {object}  types.Response
// @Router       /perm/update [POST]
// @security BearerAuth
func (r RoleHandler) UpdatePermission(ctx *gin.Context) {
	var updateOpt role.UpdatePermOption
	if err := bind.BindAndResp(ctx, bind.Json(&updateOpt)); err != nil {
		return
	}
	err := r.enforcer.UpdatePerm(updateOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.update.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.update.ok").Send()
}

// RemovePermission
// @Summary      RemovePermission
// @Description  remove the specified permission
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        id   query     types.Id  true  "perm id"
// @Success      200  {object}  types.Response
// @Router       /perm/remove [DELETE]
// @security BearerAuth
func (r RoleHandler) RemovePermission(ctx *gin.Context) {
	var permId types.Id
	if err := bind.BindAndResp(ctx, bind.Query(&permId)); err != nil {
		return
	}

	err := r.enforcer.RemovePerm(permId.Uint())
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.delete.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.delete.ok").Send()
}

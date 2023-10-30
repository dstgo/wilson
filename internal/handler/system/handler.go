package system

import (
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/auth"
	"github.com/dstgo/wilson/internal/types/code"
	"github.com/dstgo/wilson/internal/types/dict"
	"github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/dstgo/wilson/pkg/ginx/bind"
	"github.com/dstgo/wilson/pkg/ginx/httpx"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SystemProviderSet = wire.NewSet(
	NewPingHandler,
	NewPingLogic,
	NewAuthenticator,
	NewAuthHandler,
	NewRoleEnforcer,
	NewRoleHandler,
	NewAPIKey,
	NewAPIKeyHandler,
	NewDictResolver,
	NewDictHandler,
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
// @Description  [guest]
// @Description  test app api if is accessible
// @Tags         system
// @Accept       json
// @Produce      json
// @Param        name	query	system.PingRequest	true	"ping name"
// @Success      200  {object}  types.Response{data=system.PingReply}
// @Router       /ping [GET]
func (p PingHandler) Ping(ctx *gin.Context) {
	pingReq := new(system.PingRequest)
	err := bind.Binds(ctx,
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
// @Description  [guest]
// @Description  test app api authentication if is work
// @Tags         system
// @Accept       json
// @Produce      json
// @Param        name   query      system.PingRequest  true  "pong name"
// @Success      200  {object}  types.Response
// @Router       /pong [GET]
func (p PingHandler) Pong(ctx *gin.Context) {
	pongReq := new(system.PingRequest)
	err := bind.Binds(ctx,
		bind.Query(pongReq),
	)
	if err != nil {
		return
	}

	reply := p.PingLogic.Pong(pongReq.Name)

	resp.Ok(ctx).Code(code.RequestOk).Msg("ping").Data(reply).Send()
}

func NewAuthHandler(authen Authenticator) AuthHandler {
	return AuthHandler{authenticator: authen}
}

type AuthHandler struct {
	authenticator Authenticator
}

// Login
// @Summary      Login
// @Description  [guest]
// @Description  if login success, return jwt token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginBody	body	auth.LoginOption	true	"comment"
// @Success      200  {object}  types.Response{data=auth.Token}
// @Router       /auth/login [POST]
func (a AuthHandler) Login(ctx *gin.Context) {
	loginRequest := new(auth.LoginOption)
	if err := bind.Binds(ctx, bind.Json(loginRequest)); err != nil {
		return
	}
	token, err := a.authenticator.TryLogin(ctx, loginRequest.Username, loginRequest.Password, loginRequest.Persistent)
	if err != nil {
		resp.Fail(ctx).Code(code.LoginFailed).MsgI18n("auth.loginFail").Error(err).Transparent().Send()
		return
	}
	resp.Ok(ctx).Code(code.LoginOk).MsgI18n("auth.loginOk").
		Data(auth.Token{Token: token.Access.Tk.SignedJwt, Refresh: token.Refresh.Tk.SignedJwt}).Send()
}

// Register
// @Summary      Register
// @Description  [guest]
// @Description  user register api
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        registerBody	body	auth.RegisterOption	true	"comment"
// @Success      200  {object}  types.Response
// @Router       /auth/register [POST]
func (a AuthHandler) Register(ctx *gin.Context) {
	registerRequest := new(auth.RegisterOption)
	if err := bind.Binds(ctx, bind.Json(registerRequest)); err != nil {
		return
	}
	err := a.authenticator.TryRegisterNewUser(ctx, registerRequest.Username, registerRequest.Password, registerRequest.Code)
	if err != nil {
		resp.Fail(ctx).Code(code.RegisterFailed).MsgI18n("auth.registerFail").Error(err).Transparent().Send()
		return
	}
	resp.Ok(ctx).Code(code.RegisterOk).MsgI18n("auth.registerOk").Send()
}

// Logout
// @Summary      Logout
// @Description  [user]
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
	err := a.authenticator.TryLogout(ctx, tokenInfo.ID)
	if err != nil {
		resp.Fail(ctx).Code(code.LogoutFailed).MsgI18n("auth.logoutFail").Error(err).Transparent().Send()
		return
	}
	resp.Ok(ctx).Code(code.LogoutOK).MsgI18n("auth.logoutOk").Send()
}

// Refresh
// @Summary      Refresh
// @Description  [guest]
// @Description  carry refresh token in query params, access token in header
// @Description  if refresh-token expired , TokenRefresher will not refresh token [4012]
// @Description  else if access-token has expired after delay duration, TokenRefresher will not refresh token [4012]
// @Description  else if access-token has expired before delay duration, TokenRefresher will issue a new access-token [2005]
// @Description  else if access-token has not expired, TokenRefresher will renewal the 1/10 access-token ttl per time  [2005]
// @Description  else if access-token has not expired, and ttl >= 2 * conf.JwtConf.Exp, TokenRefresher will not refresh token [4013]
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        refresh   query     string  true  "refresh token"
// @Success      200  {object}  types.Response{data=auth.Token}
// @Router       /auth/refresh [GET]
func (a AuthHandler) Refresh(ctx *gin.Context) {
	var refresh auth.RefreshTokenOption
	accessToken := httpx.GetBearerTokenFromCtx(ctx)
	refresh.Access = accessToken
	if err := bind.Binds(ctx, bind.Query(&refresh)); err != nil {
		return
	}

	token, err := a.authenticator.RefreshToken(ctx, refresh.Access, refresh.Refresh)
	if err != nil {
		resp.Fail(ctx).Code(code.RefreshFailed).MsgI18n("auth.refresh.failed").Error(err).Transparent().Send()
		return
	}
	resp.Ok(ctx).Code(code.RefreshOk).MsgI18n("auth.refresh.ok").Data(auth.Token{
		Token: token.Access.Tk.SignedJwt,
	}).Send()
}

// ForgotPassword
// @Summary      ForgotPassword
// @Description  [guest]
// @Description  forgot password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        form   body      auth.ForgotPasswordOption  true  "Account ID"
// @Success      200  {object}  types.Response
// @Router       /auth/forgotpwd [POST]
func (a AuthHandler) ForgotPassword(ctx *gin.Context) {
	changePasswordReq := new(auth.ForgotPasswordOption)
	err := bind.Binds(ctx,
		bind.Json(changePasswordReq),
	)
	if err != nil {
		return
	}

	err = a.authenticator.ChangePassword(ctx, changePasswordReq.Password, changePasswordReq.Code)
	if err != nil {
		resp.Fail(ctx).MsgI18n("auth.changePasswdFail").Error(err).Transparent().Send()
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
// @Description  [admin]
// @Description  get role list by page
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        page  query    role.PageOption  true  "page option"
// @Success      200  {object}  types.Response{data=[]role.RoleInfo}
// @Router       /role/list [GET]
// @security BearerAuth
func (r RoleHandler) GetRoleList(ctx *gin.Context) {
	var pageOpt role.PageOption
	if err := bind.Binds(ctx, bind.Query(&pageOpt)); err != nil {
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
// @Description  [admin]
// @Description  get permissions list belong to role
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        queryOpt   query   system.Id true  "role perms query opt"
// @Success      200  {object}  types.Response{data=[]role.PermGroup}
// @Router       /role/perms [GET]
// @security BearerAuth
func (r RoleHandler) GetRolePerms(ctx *gin.Context) {
	var queryOpt system.Id
	if err := bind.Binds(ctx, bind.Query(&queryOpt)); err != nil {
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
// @Description  [admin]
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
	if err := bind.Binds(ctx, bind.Json(&createOpt)); err != nil {
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
// @Description  [admin]
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
	if err := bind.Binds(ctx, bind.Json(&updateOpt)); err != nil {
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
// @Description  [admin]
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
	if err := bind.Binds(ctx, bind.Json(&grantOption)); err != nil {
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
// @Description  [admin]
// @Description  remove a role,and its permission record will be deleted too
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        id   query     system.Id  true  "roleD id"
// @Success      200  {object}  types.Response
// @Router       /role/remove [DELETE]
// @security BearerAuth
func (r RoleHandler) RemoveRole(ctx *gin.Context) {
	var roleId system.Id
	if err := bind.Binds(ctx, bind.Query(&roleId)); err != nil {
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
// @Description  [admin]
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
	if err := bind.Binds(ctx, bind.Query(&pageOpt)); err != nil {
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
// @Description  [admin]
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
	if err := bind.Binds(ctx, bind.Json(&createOpt)); err != nil {
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
// @Description  [admin]
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
	if err := bind.Binds(ctx, bind.Json(&updateOpt)); err != nil {
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
// @Description  [admin]
// @Description  remove the specified permission
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        id   query     system.Id  true  "perm id"
// @Success      200  {object}  types.Response
// @Router       /perm/remove [DELETE]
// @security BearerAuth
func (r RoleHandler) RemovePermission(ctx *gin.Context) {
	var permId system.Id
	if err := bind.Binds(ctx, bind.Query(&permId)); err != nil {
		return
	}

	err := r.enforcer.RemovePerm(permId.Uint())
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.delete.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.delete.ok").Send()
}

func NewAPIKeyHandler(apikey ApiKey) APIKeyHandler {
	return APIKeyHandler{apikey: apikey}
}

type APIKeyHandler struct {
	apikey ApiKey
}

// ListAPIKeys
// @Summary      ListAPIKeys
// @Description  list specified user api keys
// @Tags         key
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response{data=[]auth.APIKey}
// @Router       /key/list [GET]
// @security BearerAuth
func (a APIKeyHandler) ListAPIKeys(ctx *gin.Context) {
	info := authen.GetContextTokenInfo(ctx)
	uuid := info.UUID
	keys, err := a.apikey.ListApiKey(ctx, uuid)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.query.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.query.ok").Data(keys).Send()
}

// CreateAPIKey
// @Summary      CreateAPIKey
// @Description  create specified user api key
// @Tags         key
// @Accept       json
// @Produce      json
// @Param        KeyCreateOption   body   auth.KeyCreateOption  true  "KeyCreateOption"
// @Success      200  {object}  types.Response
// @Router       /key/create [POST]
// @security BearerAuth
func (a APIKeyHandler) CreateAPIKey(ctx *gin.Context) {
	var createOpt auth.KeyCreateOption
	createOpt.Uid = authen.GetContextTokenInfo(ctx).UUID

	if err := bind.Binds(ctx, bind.Json(&createOpt)); err != nil {
		return
	}

	err := a.apikey.CreateAPiKey(ctx, createOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.create.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.create.ok").Send()
}

// RemoveAPIKey
// @Summary      RemoveAPIKey
// @Description  remove specified api key
// @Tags         key
// @Accept       json
// @Produce      json
// @Param        KeyRemoveOption   query     auth.KeyRemoveOption  true  "KeyRemoveOption"
// @Success      200  {object}  types.Response
// @Router       /key/remove [DELETE]
// @security BearerAuth
func (a APIKeyHandler) RemoveAPIKey(ctx *gin.Context) {
	var removeOpt auth.KeyRemoveOption
	removeOpt.UUID = authen.GetContextTokenInfo(ctx).UUID
	if err := bind.Binds(ctx, bind.Query(&removeOpt)); err != nil {
		return
	}

	err := a.apikey.RemoveApiKey(ctx, removeOpt.UUID, removeOpt.Key)
	if err != nil {
		resp.Fail(ctx).MsgI18n("op.delete.fail").Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n("op.delete.ok").Send()
}

func NewDictHandler(dt DictResolver) DictHandler {
	return DictHandler{dt: dt}
}

type DictHandler struct {
	dt DictResolver
}

// GetDictInfo
// @Summary      GetDictInfo
// @Description  get dict data info by code
// @Tags         dict
// @Accept       json
// @Produce      json
// @Param        code   query     dict.CodeOption  true  "dict id"
// @Success      200  {object}  types.Response{data=[]dict.DictDataInfo}
// @Router       /dict/info [GET]
// @security BearerAuth
func (d DictHandler) GetDictInfo(ctx *gin.Context) {
	var codeOpt dict.CodeOption
	if err := bind.Binds(ctx, bind.Query(&codeOpt)); err != nil {
		return
	}
	info, err := d.dt.GetDictInfo(ctx, codeOpt.Code)
	if err != nil {
		resp.Fail(ctx).MsgI18n(types.QueryFail).Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n(types.QueryOk).Data(info).Send()
}

// ListDict
// @Summary      ListDict
// @Description  list dict pages
// @Tags         dict
// @Accept       json
// @Produce      json
// @Param        page  query    dict.DictPageOption  true  "DictPageOption"
// @Success      200  {object}  types.Response{data=[]dict.DictDetail}
// @Router       /dict/list [GET]
// @security BearerAuth
func (d DictHandler) ListDict(ctx *gin.Context) {
	var pageOpt dict.DictPageOption
	if err := bind.Binds(ctx, bind.Query(&pageOpt)); err != nil {
		return
	}
	list, err := d.dt.ListPageDict(ctx, pageOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n(types.QueryFail).Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n(types.QueryOk).Data(list).Send()
}

// CreateDict
// @Summary      CreateDict
// @Description  create dict
// @Tags         dict
// @Accept       json
// @Produce      json
// @Param        DictSaveOption   body  dict.DictSaveOption  true  "DictSaveOption"
// @Success      200  {object}  types.Response
// @Router       /dict/create [POST]
// @security BearerAuth
func (d DictHandler) CreateDict(ctx *gin.Context) {
	var opt dict.DictSaveOption
	if err := bind.Binds(ctx, bind.Json(&opt)); err != nil {
		return
	}
	err := d.dt.CreateDict(ctx, opt)
	if err != nil {
		resp.Fail(ctx).MsgI18n(types.CreateFail).Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n(types.CreateOk).Send()
}

// UpdateDict
// @Summary      UpdateDict
// @Description  update dict
// @Tags         dict
// @Accept       json
// @Produce      json
// @Param        DictUpdateOption   body  dict.DictUpdateOption  true  "DictUpdateOption"
// @Success      200  {object}  types.Response
// @Router       /dict/update [POST]
// @security BearerAuth
func (d DictHandler) UpdateDict(ctx *gin.Context) {
	var opt dict.DictUpdateOption
	if err := bind.Binds(ctx, bind.Json(&opt)); err != nil {
		return
	}
	err := d.dt.UpdateDict(ctx, opt)
	if err != nil {
		resp.Fail(ctx).MsgI18n(types.UpdateFail).Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n(types.UpdateOk).Send()
}

// RemoveDict
// @Summary      RemoveDict
// @Description  remove dict
// @Tags         dict
// @Accept       json
// @Produce      json
// @Param        id query     system.Id  true  "dict id"
// @Success      200  {object}  types.Response
// @Router       /dict/remove [DELETE]
// @security BearerAuth
func (d DictHandler) RemoveDict(ctx *gin.Context) {
	var id system.Id
	if err := bind.Binds(ctx, bind.Query(&id)); err != nil {
		return
	}
	err := d.dt.RemoveDict(ctx, id.Uint())
	if err != nil {
		resp.Fail(ctx).MsgI18n(types.RemoveFail).Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n(types.RemoveOk).Send()
}

// ListDictData
// @Summary      ListDictData
// @Description  list dict data
// @Tags         dict
// @Accept       json
// @Produce      json
// @Param        code query     dict.DictDataPageOption  true  "dict code"
// @Success      200  {object}  types.Response{data=[]dict.DictDataDetail}
// @Router       /dict/data/list [GET]
// @security BearerAuth
func (d DictHandler) ListDictData(ctx *gin.Context) {
	var pageOpt dict.DictDataPageOption
	if err := bind.Binds(ctx, bind.Query(&pageOpt)); err != nil {
		return
	}
	list, err := d.dt.ListPageDictData(ctx, pageOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n(types.QueryFail).Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n(types.QueryOk).Data(list).Send()
}

// CreateDictData
// @Summary      CreateDictData
// @Description  get string by ID
// @Tags         dict
// @Accept       json
// @Produce      json
// @Param        DictDataSaveOption   body      dict.DictDataSaveOption  true  "DictDataSaveOption"
// @Success      200  {object}  types.Response
// @Router       /dict/data/create [POST]
// @security BearerAuth
func (d DictHandler) CreateDictData(ctx *gin.Context) {
	var saveOpt dict.DictDataSaveOption
	if err := bind.Binds(ctx, bind.Json(&saveOpt)); err != nil {
		return
	}
	err := d.dt.CreateDictData(ctx, saveOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n(types.CreateFail).Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n(types.CreateOk).Send()
}

// UpdateDictData
// @Summary      UpdateDictData
// @Description  get string by ID
// @Tags         dict
// @Accept       json
// @Produce      json
// @Param        DictDataUpdateOption   body      dict.DictDataUpdateOption  true  "DictDataUpdateOption"
// @Success      200  {object}  types.Response
// @Router       /dict/data/update [POST]
// @security BearerAuth
func (d DictHandler) UpdateDictData(ctx *gin.Context) {
	var saveOpt dict.DictDataUpdateOption
	if err := bind.Binds(ctx, bind.Json(&saveOpt)); err != nil {
		return
	}
	err := d.dt.UpdateDictData(ctx, saveOpt)
	if err != nil {
		resp.Fail(ctx).MsgI18n(types.UpdateFail).Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n(types.UpdateOk).Send()
}

// RemoveDictData
// @Summary      RemoveDictData
// @Description  get string by ID
// @Tags         dict
// @Accept       json
// @Produce      json
// @Param        system.Id  query     system.Id  true  "id"
// @Success      200  {object}  types.Response
// @Router       /dict/data/remove [DELETE]
// @security BearerAuth
func (d DictHandler) RemoveDictData(ctx *gin.Context) {
	var removeOpt system.Id
	if err := bind.Binds(ctx, bind.Query(&removeOpt)); err != nil {
		return
	}
	err := d.dt.RemoveDictData(ctx, removeOpt.Uint())
	if err != nil {
		resp.Fail(ctx).MsgI18n(types.RemoveFail).Error(err).Send()
		return
	}
	resp.Ok(ctx).MsgI18n(types.RemoveOk).Send()
}

package auth

import (
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/types/code"
	"github.com/dstgo/wilson/app/types/request"
	"github.com/dstgo/wilson/app/types/response"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var AuthProviderSet = wire.NewSet(
	NewAuthLogic,
	NewAuthHandler,
	NewRoleLogic,
	NewRoleHandler,
)

func NewAuthHandler(authLogic AuthLogic) AuthHandler {
	return AuthHandler{Authlogic: authLogic}
}

type AuthHandler struct {
	Authlogic AuthLogic
}

// Login
//
//	@Summary		user login api
//	@Description	if login success, return jwt token
//	@Tags			system/auth
//	@param			loginBody	body	request.LoginRequest	true	"comment"
//	@Accept			json
//	@Produce		json
//	@Handler			/auth/login [POST]
func (a AuthHandler) Login(ctx *gin.Context) {
	loginRequest := new(request.LoginRequest)
	if err := vax.BindAndResp(ctx, vax.Json(loginRequest)); err != nil {
		return
	}
	signedJwt, err := a.Authlogic.TryLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		resp.Fail(ctx).Code(code.LoginFailed).MsgI18n("auth.loginFail").Error(err).Send()
		return
	}
	resp.Ok(ctx).Code(code.LoginOk).MsgI18n("auth.loginOk").
		Data(response.Token{Token: signedJwt.SignedJwt}).Send()
}

// Register
//
//	@Summary		user register api
//	@Description	user register api
//	@Tags			system/auth
//	@Accept			json
//	@Produce		json
//	@Param			registerBody	body	request.RegisterRequest	true	"comment"
//	@Handler			/auth/register [POST]
func (a AuthHandler) Register(ctx *gin.Context) {
	registerRequest := new(request.RegisterRequest)
	if err := vax.BindAndResp(ctx, vax.Json(registerRequest)); err != nil {
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
//
//	@Summary		user logout api
//	@Description	user logout
//	@Tags			system/auth
//	@Produce		json
//	@Handler			/auth/logout [DELETE]
func (a AuthHandler) Logout(ctx *gin.Context) {
	// get user info from parsed request context
	tokenInfo := auth.GetContextTokenInfo(ctx)
	err := a.Authlogic.TryLogout(tokenInfo.ID)
	if err != nil {
		resp.Fail(ctx).Code(code.LogoutFailed).MsgI18n("auth.logoutFail").Error(err).Send()
		return
	}
	resp.Ok(ctx).Code(code.LogoutOK).MsgI18n("auth.logoutOk").Send()
}

func NewRoleHandler(roleLogic RoleLogic) RoleHandler {
	return RoleHandler{roleLogic}
}

type RoleHandler struct {
	RoleLogic RoleLogic
}

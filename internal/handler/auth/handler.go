package auth

import (
	"github.com/dstgo/wilson/internal/pkg/resp"
	"github.com/dstgo/wilson/internal/pkg/valid"
	"github.com/dstgo/wilson/internal/types/api/auth"
	"github.com/dstgo/wilson/internal/types/code"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var AuthProviderSet = wire.NewSet(
	TokenCacheProviderSet,
	NewAuthenticator,
	NewAuthHandler,
)

func NewAuthHandler(authen Authenticator) AuthHandler {
	return AuthHandler{Authlogic: authen}
}

type AuthHandler struct {
	Authlogic Authenticator
}

// Login
//
//	@Summary		Login
//	@Description	if login success, return jwt token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@param			loginBody	body	auth.LoginOption	true	"comment"
//	@Router			/auth/login [POST]
func (a AuthHandler) Login(ctx *gin.Context) {
	loginRequest := new(auth.LoginOption)
	if err := valid.BindAndResp(ctx, valid.Json(loginRequest)); err != nil {
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
//
//	@Summary		Register
//	@Description	user register api
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			registerBody	body	auth.RegisterOption	true	"comment"
//	@Router			/auth/register [POST]
func (a AuthHandler) Register(ctx *gin.Context) {
	registerRequest := new(auth.RegisterOption)
	if err := valid.BindAndResp(ctx, valid.Json(registerRequest)); err != nil {
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
//	@Summary		Logout
//	@Description	user logout
//	@Tags			auth
//	@Produce		json
//	@Router			/auth/logout [DELETE]
func (a AuthHandler) Logout(ctx *gin.Context) {
	// get user info from parsed request context
	tokenInfo := GetContextTokenInfo(ctx)
	err := a.Authlogic.TryLogout(tokenInfo.ID)
	if err != nil {
		resp.Fail(ctx).Code(code.LogoutFailed).MsgI18n("auth.logoutFail").Error(err).Send()
		return
	}
	resp.Ok(ctx).Code(code.LogoutOK).MsgI18n("auth.logoutOk").Send()
}

// ForgotPassword
//
//	@Summary		ForgotPasswd
//	@Description	forgot password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			changePasswdBoy	body	auth.ForgotPasswordOption	true	"comment"
//	@Router			/auth/forgotpwd [POST]
func (a AuthHandler) ForgotPassword(ctx *gin.Context) {
	changePasswordReq := new(auth.ForgotPasswordOption)
	err := valid.BindAndResp(ctx,
		valid.Json(changePasswordReq),
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

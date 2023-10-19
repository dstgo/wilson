package system

import (
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/core/bind"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/data/cache"
	"github.com/dstgo/wilson/internal/types/api/auth"
	"github.com/dstgo/wilson/internal/types/api/system"
	"github.com/dstgo/wilson/internal/types/code"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SystemProviderSet = wire.NewSet(
	cache.TokenCacheProvider,
	NewPingHandler,
	NewPingLogic,
	NewAuthenticator,
	NewAuthHandler,
)

func NewPingHandler(logic PingLogic) PingHandler {
	return PingHandler{
		PingLogic: logic,
	}
}

type PingHandler struct {
	PingLogic PingLogic
}

// Ping
// @Summary      Ping
// @Description  test app api if is accessible
// @Tags         system
// @Accept       json
// @Produce      json
// @Param        name	query	system.PingRequest	true	"ping name"
// @Success      200  {object}  api.Response{data=auth.PingReply}
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
// @Success      200  {object}  api.Response
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
// @Success      200  {object}  api.Response{data=auth.Token}
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
// @Success      200  {object}  api.Response
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
// @Success      200  {object}  api.Response
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
// @Success      200  {object}  api.Response
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

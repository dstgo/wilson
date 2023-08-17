package system

import (
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/logic/systemLogic"
	"github.com/dstgo/wilson/app/pkg/errorx"
	"github.com/dstgo/wilson/app/types/request"
	"github.com/dstgo/wilson/app/types/response"
	"github.com/gin-gonic/gin"
)

func NewAuthApi(authLogic systemLogic.AuthLogic) AuthApi {
	return AuthApi{authlogic: authLogic}
}

type AuthApi struct {
	authlogic systemLogic.AuthLogic
}

// Login
//
//	@Summary		user login api
//	@Description	if login success, return jwt token
//	@Tags			system/auth
//	@param			name	body	request.LoginRequest	true	"comment"
//	@Accept			json
//	@Produce		json
//	@Router			/auth/login [POST]
func (a AuthApi) Login(ctx *gin.Context) {
	loginRequest := new(request.LoginRequest)
	if err := vax.BindsAndResp(ctx, vax.Json(loginRequest)); err != nil {
		return
	}
	signedJwt, err := a.authlogic.TryLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		resp.Fail(ctx, 4001, errorx.WrapI18n(ctx, err, "app.loginFailed"))
	}
	resp.OkI18n(ctx, 2001, "app.loginSuccess", response.Token{Token: signedJwt.SignedJwt})
}

// Register
//
//	@Summary		user register api
//	@Description	user register api
//	@Tags			system/auth
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	true	"comment"
//	@Success		200		{object}	string	"comment"
//	@Failure		400		{object}	string	"comment"
//	@Router			/auth/register [POST]
func (a AuthApi) Register(ctx *gin.Context) {

}

func (a AuthApi) Logout(ctx *gin.Context) {

}

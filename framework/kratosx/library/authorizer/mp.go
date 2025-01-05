package authorizer

import (
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/dstgo/wilson/framework/kratosx"
)

type wx struct {
}

func init() {
	register("mp", &wx{})
}

func (w wx) Name() string {
	return "微信"
}

func (w wx) GetAccessToken(ctx kratosx.Context, req GetAccessTokenRequest) (*GetAccessTokenReply, error) {
	data := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}{}

	response, err := ctx.Http().Option(func(request *resty.Request) {
		request.SetQueryParam("appid", req.Ak).
			SetQueryParam("secret", req.Sk).
			SetQueryParam("grant_type", "client_credential")
	}).Post("https://api.weixin.qq.com/cgi-bin/token", nil)
	if err != nil {
		return nil, err
	}

	if err = response.Result(&data); err != nil {
		return nil, err
	}

	return &GetAccessTokenReply{
		Token:  data.AccessToken,
		Expire: time.Now().Unix() + int64(data.ExpiresIn),
	}, nil
}

func (w wx) GetAuthInfo(ctx kratosx.Context, req GetAuthInfoRequest) (*GetAuthInfoReply, error) {
	data := struct {
		Openid  string `json:"openid"`
		Unionid string `json:"unionid"`
	}{}
	response, err := ctx.Http().Option(func(request *resty.Request) {
		request.SetQueryParam("appid", req.Ak).
			SetQueryParam("secret", req.Sk).
			SetQueryParam("js_code", req.Code).
			SetQueryParam("grant_type", "authorization_code")
	}).Post("https://api.weixin.qq.com/sns/jscode2session", nil)
	if err != nil {
		return nil, err
	}
	if err = response.Result(&data); err != nil {
		return nil, err
	}
	return &GetAuthInfoReply{
		AuthId:  data.Openid,
		UnionId: &data.Unionid,
	}, nil
}

package authorizer

import (
	"errors"
	"time"

	json "github.com/json-iterator/go"

	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/pkg/lock"
)

type Authorizer interface {
	Name() string
	GetAccessToken(kratosx.Context, GetAccessTokenRequest) (*GetAccessTokenReply, error)
	GetAuthInfo(kratosx.Context, GetAuthInfoRequest) (*GetAuthInfoReply, error)
}

type GetAuthInfoRequest struct {
	Ak    string
	Sk    string
	Extra string
	Code  string
	Token string
}

type GetAuthInfoReply struct {
	AuthId  string
	UnionId *string
}

type GetAccessTokenRequest struct {
	Ak    string
	Sk    string
	Extra string
	Code  string
}

type GetAccessTokenReply struct {
	Token  string
	Expire int64
}

type Interface interface {
	GetAuthorizers() map[string]Authorizer
	GetToken(ctx kratosx.Context) (*GetAccessTokenReply, error)
	GetAuthInfo(ctx kratosx.Context, token string) (*GetAuthInfoReply, error)
}

type authorizer struct {
	config *Config
}

type Config struct {
	Ak       string
	Sk       string
	Code     string
	Extra    string
	Platform string
}

var ins = make(map[string]Authorizer)

func register(key string, at Authorizer) {
	ins[key] = at
}

func New(config *Config) Interface {
	return &authorizer{
		config: config,
	}
}

func (ath *authorizer) GetAuthorizers() map[string]Authorizer {
	return ins
}

func (ath *authorizer) GetAuthorizer(platform string) Authorizer {
	return ins[platform]
}

func (ath *authorizer) GetToken(ctx kratosx.Context) (*GetAccessTokenReply, error) {
	var (
		reply   *GetAccessTokenReply
		key     = "auth:token:" + ath.config.Platform
		lockKey = "auth:token:" + ath.config.Platform + ":lock"
	)

	lk := lock.New(ctx.Redis(), lockKey)
	err := lk.AcquireFunc(ctx, func() error {
		err := ctx.Redis().Get(ctx, key).Scan(reply)
		return err
	}, func() error {
		atr := ath.GetAuthorizer(ath.config.Platform)
		if atr == nil {
			return errors.New("not exist channel authorizer")
		}
		res, err := atr.GetAccessToken(ctx, GetAccessTokenRequest{
			Ak:    ath.config.Ak,
			Sk:    ath.config.Sk,
			Code:  ath.config.Code,
			Extra: ath.config.Extra,
		})
		if err != nil {
			return err
		}
		resStr, _ := json.MarshalToString(res)
		ctx.Redis().Set(ctx, key, resStr, time.Duration(res.Expire-time.Now().Unix()-10)*time.Second)
		reply = res
		return nil
	})

	return reply, err
}

func (ath *authorizer) GetAuthInfo(ctx kratosx.Context, token string) (*GetAuthInfoReply, error) {
	atr := ath.GetAuthorizer(ath.config.Platform)
	if atr == nil {
		return nil, errors.New("not exist channel authorizer")
	}
	return atr.GetAuthInfo(ctx, GetAuthInfoRequest{
		Ak:    ath.config.Ak,
		Sk:    ath.config.Sk,
		Code:  ath.config.Code,
		Extra: ath.config.Extra,
		Token: token,
	})
}

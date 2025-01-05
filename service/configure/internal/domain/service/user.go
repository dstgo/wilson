package service

import (
	"encoding/base64"
	"time"

	"github.com/forgoer/openssl"
	json "github.com/json-iterator/go"

	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/api/gen/errors"

	"github.com/dstgo/wilson/service/configure/internal/conf"
)

type User struct {
	conf *conf.Config
}

func NewUser(conf *conf.Config) *User {
	return &User{
		conf: conf,
	}
}

type Password struct {
	Password string `json:"password"`
	Time     int64  `json:"time"`
}

// Login 用户登录
func (us *User) Login(ctx kratosx.Context, username, password string) (string, error) {
	// 密码解密
	pwByte, _ := base64.StdEncoding.DecodeString(password)
	decryptData, err := openssl.RSADecrypt(pwByte, ctx.Loader("login"))
	if err != nil {
		return "", errors.PasswordError(err.Error())
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return "", errors.PasswordError()
	}

	// 判断当前时间戳是否过期,超过10s则拒绝
	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		return "", errors.PasswordExpireError()
	}

	// 判断用户密码是否正确
	if us.conf.Author.AdminPassword != pw.Password || us.conf.Author.AdminUser != username {
		return "", errors.PasswordError()
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(map[string]any{"userId": 1, "roleKeyword": "superAdmin"})
	if err != nil {
		return "", errors.SystemError(err.Error())
	}

	return token, nil
}

// RefreshToken 刷新用户token
func (us *User) RefreshToken(ctx kratosx.Context) (string, error) {
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return "", errors.RefreshTokenError(err.Error())
	}
	return token, nil
}

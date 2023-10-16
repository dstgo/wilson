package auth

import (
	"context"
	"errors"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/handler/email"
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/dstgo/wilson/internal/pkg/jwtx"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/pkg/resp"
	"github.com/dstgo/wilson/pkg/vax/is"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func NewAuthenticator(cfg *conf.AppConf, userData user.InfoData, codeCache email.CodeCache, tokenCache TokenCache) Authenticator {
	return Authenticator{
		issue:      NewCacheAuthor(cfg.JwtConf, tokenCache),
		userData:   userData,
		codeCache:  codeCache,
		tokenCache: tokenCache,
	}
}

type Authenticator struct {
	issue      Issuer
	userData   user.InfoData
	codeCache  email.CodeCache
	tokenCache TokenCache
}

func (a Authenticator) TryLogin(userName string, password string) (jwtx.Jwt, error) {
	var token jwtx.Jwt

	var (
		user    entity.User
		userErr error
	)

	// try to find the user
	if err := is.Email.Validate(locale.L().Default(), userName); err != nil {
		user, userErr = a.userData.GetUserByEmail(userName)
	} else {
		user, userErr = a.userData.GetUserByName(userName)
	}

	// if user not found, return error
	if errors.Is(userErr, gorm.ErrRecordNotFound) {
		return token, resp.NewErr().Status(http.StatusNotFound).I18n("user.notfound")
	} else if userErr != nil {
		return token, resp.DataBaseErr(userErr)
	}

	// compare the password
	sum := cryptor.Sha512WithBase64(password)
	if sum != user.Password {
		return token, resp.NewErr().Status(http.StatusBadRequest).I18n("user.wrongPassword")
	}

	// issue token
	issueToken, err := a.issue.Issue(context.Background(), UserPayload{
		Username: user.Username,
		UserID:   user.UUID,
	}, -1)

	if err != nil {
		return token, resp.ProgramErr(err)
	}

	token = issueToken
	return token, nil
}

func (a Authenticator) TryRegisterNewUser(username string, password string, code string) error {
	ctx := context.Background()
	// find the authcode from redis
	cacheEmail, err := a.codeCache.Check(ctx, code)
	if errors.Is(err, redis.Nil) {
		return resp.NewErr().Status(http.StatusBadRequest).I18n("email.codeExpired")
	} else if err != nil {
		return resp.DataBaseErr(err)
	}

	// try to find the userInfo
	userInfo, err := a.userData.GetUserByName(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp.DataBaseErr(err)
	}

	// if userInfo found, return error
	if len(userInfo.Username) > 0 {
		return resp.NewErr().Status(http.StatusBadRequest).I18n("userInfo.alreadyExist")
	}

	// encrypt
	password = cryptor.Sha512WithBase64(password)

	// create new userInfo
	newUser := entity.User{
		UUID:     uuid.NewString(),
		Username: username,
		Password: password,
		Email:    cacheEmail,
	}

	if err := a.userData.CreateUser(newUser); err != nil {
		return resp.DataBaseErr(err)
	}

	return nil
}

func (a Authenticator) TryLogout(tokenId string) error {
	ctx := context.Background()

	_, ok, err := a.tokenCache.Get(ctx, tokenId)
	// check if token is expired
	if !ok && err == nil {
		return resp.NewErr().Status(http.StatusUnauthorized).I18n("jwt.expired")
	} else if err != nil {
		return resp.DataBaseErr(err)
	}

	err = a.tokenCache.Del(ctx, tokenId)
	if err != nil {
		return resp.DataBaseErr(err)
	}

	return nil
}

func (a Authenticator) ChangePassword(newPassword string, code string) error {
	ctx := context.Background()

	// get email
	emailCache, err := a.codeCache.Check(ctx, code)
	if errors.Is(err, redis.Nil) {
		return resp.NewErr().Status(http.StatusBadRequest).I18n("email.codeExpired")
	} else if err != nil {
		return resp.DataBaseErr(err)
	}

	// find user by email
	userInfo, err := a.userData.GetUserByEmail(emailCache)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return resp.NewErr().Status(http.StatusNotFound).I18n("user.notfound")
	} else if err != nil {
		return resp.DataBaseErr(err)
	}

	// change the password
	userInfo.Password = cryptor.Sha512WithBase64(newPassword)

	// save
	if err := a.userData.UpdateUserInfo(userInfo); err != nil {
		return resp.DataBaseErr(err)
	}

	return nil
}

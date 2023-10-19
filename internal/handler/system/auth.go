package system

import (
	"context"
	"errors"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/cache"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/dstgo/wilson/internal/pkg/jwtx"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/dstgo/wilson/pkg/vax/is"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func NewAuthenticator(cfg *conf.AppConf, ds *data.DataSource, userData user.UserData, codeCache cache.RedisEmailCodeCache, tokenCache cache.TokenCache) Authenticator {
	return Authenticator{
		issue:      authen.NewCacheAuthor(cfg.JwtConf, tokenCache),
		userData:   userData,
		codeCache:  codeCache,
		tokenCache: tokenCache,
		ds:         ds,
	}
}

type Authenticator struct {
	issue      authen.Issuer
	userData   user.UserData
	codeCache  cache.RedisEmailCodeCache
	tokenCache cache.TokenCache
	ds         *data.DataSource
}

func (a Authenticator) TryLogin(userName string, password string) (jwtx.Jwt, error) {
	var token jwtx.Jwt

	var (
		userEntity entity.User
		userErr    error
	)

	// try to find the user
	if err := is.EmailFormat.Validate(locale.L().Default(), userName); err != nil {
		userEntity, userErr = a.userData.GetUserByName(a.ds.ORM(), userName)
	} else {
		userEntity, userErr = a.userData.GetUserByEmail(a.ds.ORM(), userName)
	}

	// if user not found, return error
	if errors.Is(userErr, gorm.ErrRecordNotFound) {
		return token, errs.NewError().Status(http.StatusNotFound).I18n("user.notfound")
	} else if userErr != nil {
		return token, errs.DataBaseErr(userErr)
	}

	// compare the password
	sum := cryptor.Sha512WithBase64(password)
	if sum != userEntity.Password {
		return token, errs.NewError().Status(http.StatusBadRequest).I18n("user.wrongPassword")
	}

	// issue token
	issueToken, err := a.issue.Issue(context.Background(), authen.UserPayload{
		Username: userEntity.Username,
		UUID:     userEntity.UUID,
	}, -1)

	if err != nil {
		return token, errs.ProgramErr(err)
	}

	token = issueToken
	return token, nil
}

func (a Authenticator) TryRegisterNewUser(username string, password string, code string) error {
	ctx := context.Background()
	// find the authcode from redis
	cacheEmail, err := a.codeCache.Check(ctx, code)
	if errors.Is(err, redis.Nil) {
		return errs.BadRequest(err).I18n("email.codeExpired")
	} else if err != nil {
		return errs.DataBaseErr(err)
	}

	// try to find the userInfo
	userInfo, err := a.userData.GetUserByName(a.ds.ORM(), username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.DataBaseErr(err)
	}

	// if userInfo found, return error
	if len(userInfo.Username) > 0 {
		return errs.BadRequest(err).I18n("userInfo.alreadyExist")
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

	if err := a.userData.CreateUser(a.ds.ORM(), newUser); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}

func (a Authenticator) TryLogout(tokenId string) error {
	ctx := context.Background()

	_, ok, err := a.tokenCache.Get(ctx, tokenId)
	// check if token is expired
	if !ok && err == nil {
		return errs.UnAuthorized(err).I18n("jwt.expired")
	} else if err != nil {
		return errs.DataBaseErr(err)
	}

	err = a.tokenCache.Del(ctx, tokenId)
	if err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}

func (a Authenticator) ChangePassword(newPassword string, code string) error {
	ctx := context.Background()

	// get email
	emailCache, err := a.codeCache.Check(ctx, code)
	if errors.Is(err, redis.Nil) {
		return errs.BadRequest(err).I18n("email.codeExpired")
	} else if err != nil {
		return errs.DataBaseErr(err)
	}

	// find user by email
	userInfo, err := a.userData.GetUserByEmail(a.ds.ORM(), emailCache)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ResourceNotFound(err).I18n("user.notfound")
	} else if err != nil {
		return errs.DataBaseErr(err)
	}

	// change the password
	userInfo.Password = cryptor.Sha512WithBase64(newPassword)

	// save
	if err := a.userData.UpdateUserInfo(a.ds.ORM(), userInfo); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}

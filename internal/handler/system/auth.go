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
	"github.com/dstgo/wilson/internal/types/auth"
	emailType "github.com/dstgo/wilson/internal/types/email"
	roleType "github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
	usert "github.com/dstgo/wilson/internal/types/user"
	"github.com/dstgo/wilson/pkg/vax/is"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"net/http"
)

func NewAuthenticator(cfg *conf.AppConf, ds *data.DataSource, codeCache cache.RedisEmailCodeCache, tokenCache cache.TokenCache) Authenticator {
	info := user.NewUserInfo(ds)
	modify := user.NewUserModify(ds, info)
	return Authenticator{
		issue:      authen.NewCacheAuthor(cfg.JwtConf, tokenCache),
		codeCache:  codeCache,
		tokenCache: tokenCache,
		ds:         ds,
		userInfo:   info,
		userModify: modify,
	}
}

type Authenticator struct {
	userInfo   user.UserInfo
	userModify user.UserModify
	issue      authen.Issuer
	codeCache  cache.RedisEmailCodeCache
	tokenCache cache.TokenCache
	ds         *data.DataSource
}

func (a Authenticator) TryLogin(ctx context.Context, userName string, password string) (jwtx.Jwt, error) {
	var token jwtx.Jwt

	var (
		userEntity entity.User
		userErr    error
	)

	// try to find the user
	if err := is.EmailFormat.Validate(locale.L().Default(), userName); err != nil {
		userEntity, userErr = user.GetUserByName(a.ds.ORM(), userName)
	} else {
		userEntity, userErr = user.GetUserByEmail(a.ds.ORM(), userName)
	}

	// if user not found, return error
	if errors.Is(userErr, gorm.ErrRecordNotFound) {
		return token, usert.ErrUserNotFound
	} else if userErr != nil {
		return token, system.ErrDatabase.Wrap(userErr)
	}

	// compare the password
	sum := cryptor.Sha512WithBase64(password)
	if sum != userEntity.Password {
		return token, auth.ErrWrongPassword
	}

	// issue token
	issueToken, err := a.issue.Issue(ctx, authen.UserPayload{
		Username: userEntity.Username,
		UUID:     userEntity.UUID,
	}, -1)

	if err != nil {
		return token, system.ErrProgram.Wrap(err)
	}

	token = issueToken
	return token, nil
}

func (a Authenticator) TryRegisterNewUser(ctx context.Context, username string, password string, code string) error {
	// find the authcode from redis
	cacheEmail, err := a.codeCache.Check(ctx, code)
	if errors.Is(err, redis.Nil) {
		return emailType.ErrCodeExpired
	} else if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	// find by email
	findByEmail, err := a.userInfo.GetUserInfoByEmail(cacheEmail)
	if err != nil && !errors.Is(err, usert.ErrUserNotFound) {
		return err
	} else if findByEmail.UUID != "" {
		return usert.ErrEmailAlreadyUsed
	}

	// find user by username
	findUser, err := a.userInfo.GetUserInfoByName(username)
	if err != nil && !errors.Is(err, usert.ErrUserNotFound) {
		return err
	} else if findUser.UUID != "" {
		return usert.ErrUsernameAlreadyUsed
	}

	// create new user
	createUserOption := usert.CreateUserOption{
		Username: username,
		Email:    cacheEmail,
		Password: password,
		// default user role
		Roles: []string{roleType.UserRole.Code},
	}

	if err = a.userModify.Create(createUserOption); err != nil {
		return err
	}

	return nil
}

func (a Authenticator) TryLogout(ctx context.Context, tokenId string) error {
	_, ok, err := a.tokenCache.Get(ctx, tokenId)
	// check again if token is expired
	if !ok && err == nil {
		return auth.ErrJwtExpired.Wrap(err).Status(http.StatusUnauthorized)
	} else if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	err = a.tokenCache.Del(ctx, tokenId)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	return nil
}

func (a Authenticator) ChangePassword(ctx context.Context, newPassword string, code string) error {
	// get email
	emailCache, err := a.codeCache.Check(ctx, code)
	if errors.Is(err, redis.Nil) {
		return emailType.ErrCodeExpired
	} else if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	// find user by email
	userInfo, err := user.GetUserByEmail(a.ds.ORM(), emailCache)
	if errors.Is(err, gorm.ErrRecordNotFound) || err == nil && userInfo.Id == 0 {
		return usert.ErrUserNotFound
	} else if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	// change the password
	userInfo.Password = cryptor.Sha512WithBase64(newPassword)

	// save
	if err := user.UpdateUserInfo(a.ds.ORM(), userInfo); err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	return nil
}

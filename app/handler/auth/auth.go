package auth

import (
	"context"
	"errors"
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/data/entity"
	"github.com/dstgo/wilson/app/handler/email"
	"github.com/dstgo/wilson/app/handler/user"
	"github.com/dstgo/wilson/app/pkg/jwtx"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func NewAuthLogic(issue auth.Issuer, userData user.InfoData, codeCache email.CodeCache, tokenCache TokenCache) AuthLogic {
	return AuthLogic{
		issue:      issue,
		userData:   userData,
		codeCache:  codeCache,
		tokenCache: tokenCache,
	}
}

type AuthLogic struct {
	userData   user.InfoData
	issue      auth.Issuer
	codeCache  email.CodeCache
	tokenCache TokenCache
}

func (a AuthLogic) TryLogin(userName string, password string) (jwtx.Jwt, error) {
	var token jwtx.Jwt
	// try to find the user
	user, err := a.userData.GetUserByName(userName)
	// if user not found, return error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return token, resp.NewErr().Status(http.StatusNotFound).I18n("user.notfound")
	} else if err != nil {
		return token, resp.DataBaseErr(err)
	}

	// compare the password
	sum := cryptor.Sha512WithBase64(password)
	if sum != user.Password {
		return token, resp.NewErr().Status(http.StatusBadRequest).I18n("user.wrongPassword")
	}

	// issue token
	issueToken, err := a.issue.Issue(context.Background(), auth.UserPayload{
		Username: user.Username,
		UserId:   user.UUID,
	}, -1)

	if err != nil {
		return token, resp.ProgramErr(err)
	}

	token = issueToken
	return token, nil
}

func (a AuthLogic) TryRegisterNewUser(username string, password string, code string) error {
	ctx := context.Background()
	// find the authcode from redis
	cacheEmail, err := a.codeCache.Get(ctx, code)
	if errors.Is(err, redis.Nil) {
		return resp.NewErr().Status(http.StatusBadRequest).I18n("cacheEmail.codeExpired")
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
		UUID:      uuid.NewString(),
		Username:  username,
		Password:  password,
		Email:     cacheEmail,
		CreatedAt: time.Now(),
	}

	if err := a.userData.CreateUser(newUser); err != nil {
		return resp.DataBaseErr(err)
	}

	return nil
}

func (a AuthLogic) TryLogout(tokenId string) error {
	ctx := context.Background()

	_, err := a.tokenCache.Get(ctx, tokenId)
	// check if token is expired
	if errors.Is(err, redis.Nil) {
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

//func (a AuthLogic) ChangePassword(password string, code string) error {
//
//}

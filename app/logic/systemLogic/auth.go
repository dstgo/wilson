package systemLogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/pkg/jwtx"
	"github.com/dstgo/wilson/app/repo/dao/userDao"
	"github.com/dstgo/wilson/app/repo/data"
	"github.com/dstgo/wilson/app/repo/data/entity"
	"github.com/dstgo/wilson/app/types/errs"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func NewAuthLogic(issue auth.Issuer, userDao userDao.UserInfoDao, datasource *data.DataSource) AuthLogic {
	return AuthLogic{
		issue:   issue,
		UserDao: userDao,
		redis:   datasource.Redis(),
	}
}

type AuthLogic struct {
	UserDao userDao.UserInfoDao
	issue   auth.Issuer
	redis   *redis.Client
}

func (a AuthLogic) TryLogin(userName string, password string) (jwtx.Jwt, error) {
	var token jwtx.Jwt
	// try to find the user
	user, err := a.UserDao.GetUserByName(userName)
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

	// find the authcode from redis
	email, err := a.redis.Get(context.Background(), fmt.Sprintf("email:code:%s", code)).Result()
	if errors.Is(err, redis.Nil) {
		return resp.NewErr().Status(http.StatusBadRequest).I18n("email.codeExpired")
	} else if err != nil {
		return resp.DataBaseErr(err)
	}

	// try to find the user
	user, err := a.UserDao.GetUserByName(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp.DataBaseErr(err)
	}

	// if user found, return error
	if len(user.Username) > 0 {
		return resp.NewErr().Status(http.StatusBadRequest).I18n("user.alreadyExist")
	}

	// encrypt
	password = cryptor.Sha512WithBase64(password)

	// create new user
	newUser := entity.User{
		UUID:     uuid.NewString(),
		Username: username,
		Password: password,
		Email:    email,
		CreateAt: time.Now(),
	}

	if err := a.UserDao.CreateUser(newUser); err != nil {
		return resp.DataBaseErr(err)
	}

	return nil
}

func (a AuthLogic) TryLogout(tokenId string) error {
	key := fmt.Sprintf("token:%s", tokenId)

	_, err := a.redis.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return resp.NewErr().Status(http.StatusUnauthorized).I18n("jwt.expired")
	} else if err != nil {
		return resp.DataBaseErr(err)
	}

	impact, err := a.redis.Del(context.Background(), key).Result()
	if err != nil {
		return resp.DataBaseErr(err)
	} else if impact == 0 {
		return resp.ProgramErr(errs.ErrInvalidDatabaseOperation)
	}

	return nil
}

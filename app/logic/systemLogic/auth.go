package systemLogic

import (
	"context"
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/dao/userDao"
	"github.com/dstgo/wilson/app/pkg/jwtx"
	"github.com/duke-git/lancet/v2/cryptor"
)

func NewAuthLogic(issue auth.Issuer, userDao userDao.UserInfoDao) AuthLogic {
	return AuthLogic{
		issue:   issue,
		userDao: userDao,
	}
}

type AuthLogic struct {
	issue   auth.Issuer
	userDao userDao.UserInfoDao
}

func (a AuthLogic) TryLogin(userName string, password string) (jwtx.Jwt, error) {
	var token jwtx.Jwt
	// try to find the user
	user, err := a.userDao.GetUserByName(userName)
	if err != nil {
		return token, resp.NewErr(500, err)
	}

	// if user not found, return error
	if len(user.Username) == 0 {
		return token, resp.NewI18nErr(404, "user.notfound")
	}

	// compare the password
	sum := cryptor.Sha512WithBase64(password)
	if sum != user.Password {
		return token, resp.NewI18nErr(400, "user.errPassword")
	}

	// issue token
	issueToken, err := a.issue.Issue(context.Background(), auth.UserPayload{
		Username: user.Username,
		UserId:   user.UUID,
	}, -1)

	if err != nil {
		return token, resp.NewI18nErr(500, "jwt.issueFailed")
	}

	token = issueToken

	return token, nil
}

func (a AuthLogic) Register() {

}

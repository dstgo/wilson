package user

import "github.com/dstgo/wilson/internal/handler/user"

func NewInfoLogic(userData user.UserData) InfoLogic {
	return InfoLogic{userData: userData}
}

type InfoLogic struct {
	userData user.UserData
}

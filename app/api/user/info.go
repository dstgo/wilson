package user

import "github.com/dstgo/wilson/app/handler/user"

func NewInfoLogic(userData user.InfoData) InfoLogic {
	return InfoLogic{userData: userData}
}

type InfoLogic struct {
	userData user.InfoData
}

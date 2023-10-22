package user

import (
	"github.com/dstgo/wilson/internal/data"
)

func NewInfoLogic(ds *data.DataSource) InfoLogic {
	return InfoLogic{ds: ds}
}

type InfoLogic struct {
	ds *data.DataSource
}

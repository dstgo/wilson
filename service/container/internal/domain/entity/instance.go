package entity

import (
	"github.com/dstgo/wilson/framework/pkg/sqlx"
)

type Instance struct {
	ID          int64                       `json:"id"`
	UserID      int64                       `json:"userId"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Metadata    sqlx.Json[InstanceMetaData] `json:"metadata"`
	CreatedAt   int64                       `json:"createdAt"`
	UpdatedAt   int64                       `json:"updatedAt"`
}

type InstanceMetaData struct {
	// TODO
}

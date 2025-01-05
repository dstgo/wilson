package entity

import "github.com/dstgo/wilson/framework/kratosx/types"

type Template struct {
	types.BaseModel
	ServerId    uint32 `json:"serverId" gorm:"column:server_id"`
	Content     string `json:"content" gorm:"column:content"`
	Version     string `json:"version" gorm:"column:version"`
	IsUse       bool   `json:"isUse" gorm:"column:is_use"`
	Format      string `json:"format" gorm:"column:format"`
	Description string `json:"description" gorm:"column:description"`
	Compare     string `json:"compare" gorm:"column:compare"`
}

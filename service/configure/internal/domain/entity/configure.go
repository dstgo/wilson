package entity

import "github.com/dstgo/wilson/framework/kratosx/types"

type Configure struct {
	types.BaseModel
	ServerId    uint32  `json:"serverId" gorm:"server_id"`
	EnvId       uint32  `json:"envId" gorm:"env_id"`
	Content     string  `json:"content" gorm:"content"`
	Version     string  `json:"version" gorm:"version"`
	Format      string  `json:"format" gorm:"format"`
	Description *string `json:"description" gorm:"description"`
}

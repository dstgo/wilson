package entity

import "github.com/dstgo/wilson/framework/kratosx/types"

type Resource struct {
	Keyword         string            `json:"keyword" gorm:"column:keyword"`
	Fields          string            `json:"fields" gorm:"column:fields"`
	Tag             string            `json:"tag" gorm:"column:tag"`
	Private         *bool             `json:"private" gorm:"column:private"`
	Description     *string           `json:"description" gorm:"column:description"`
	Servers         []*Server         `json:"servers" gorm:"many2many:resource_server"`
	ResourceServers []*ResourceServer `json:"resourceServers"`
	types.BaseModel
}

type ResourceServer struct {
	ServerId   uint32 `json:"serverId"  gorm:"column:server_id"`
	ResourceId uint32 `json:"resourceId"  gorm:"column:resource_id"`
}

type ResourceValue struct {
	EnvId      uint32    `json:"envId" gorm:"column:env_id"`
	ResourceId uint32    `json:"resourceId" gorm:"column:resource_id"`
	Value      string    `json:"value" gorm:"column:value"`
	Resource   *Resource `json:"resource"`
	types.BaseModel
}

package entity

import "github.com/dstgo/wilson/framework/kratosx/types"

type Server struct {
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *bool   `json:"status" gorm:"column:status"`
	types.BaseModel
}

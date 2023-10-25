package entity

import (
	"fmt"
	"gorm.io/gorm"
)

// Node
// represents a physical machine, which runs a lot of containers
type Node struct {
	gorm.Model
	Id      uint   `gorm:"primarykey"`
	Name    string `gorm:"type:varchar(50);comment:node name;"`
	Address string `gorm:"type:varchar(30);uniqueIndex;comment:node address;"`
	Note    string `gorm:"type:varchar(100);comment:node note;"`

	CreatedAt uint64 `gorm:"autoCreateTime:nano;"`
	UpdatedAt uint64 `gorm:"autoUpdateTime:nano;"`

	Instances []Instance `gorm:"foreignKey:NodeId;"`
	NodeTable
}

type NodeTable struct{}

func (n NodeTable) BeforeCreate(db *gorm.DB) error {
	return db.Set(tableOptions, fmt.Sprintf("comment '%s'", n.TableComment())).Error
}

func (n NodeTable) TableName() string {
	return "nodes"
}

func (n NodeTable) TableComment() string {
	return "remote nodes record table"
}

package entity

import "gorm.io/gorm"

// Node
// represents a physical machine, which runs a lot of containers
type Node struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);comment:node name;"`
	Address string `gorm:"type:varchar(255);comment:node address;"`
	Note    string `gorm:"type:varchar(255);comment:node note;"`

	NodeTable
}

type NodeTable struct{}

func (n NodeTable) TableName() string {
	return "nodes"
}

func (n NodeTable) TableComment() string {
	return "remote nodes record table"
}

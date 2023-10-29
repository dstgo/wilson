package entity

import (
	"fmt"
	"gorm.io/gorm"
)

// Instance represents an instance, usually a docker container
type Instance struct {
	Id     uint   `gorm:"primaryKey;"`
	Uid    string `gorm:"type:varchar(100);uniqueIndex;comment:safe unique id, sha1 from instance.id;"`
	Name   string `gorm:"type:varchar(50);comment:docker container name;"`
	Image  string `gorm:"type:varchar(100);comment:docker image label;"`
	Note   string `gorm:"type:varchar(100);comment:remark note;"`
	Cpu    uint64 `gorm:"comment:cpu count limit;"`
	Memory uint64 `gorm:"comment:memory limit;"`
	Disk   uint64 `gorm:"comment:disk limit;"`

	CreatedAt uint64 `gorm:"autoCreateTime:nano;"`
	UpdatedAt uint64 `gorm:"autoUpdateTime:nano;"`

	// foreign keys
	UserId uint `gorm:"comment:id of user who own instance;"`
	User   User `gorm:"foreignKey:UserId;"`

	NodeId uint `gorm:"comment:id of node which own instance;"`
	Node   Node `gorm:"foreignKey:NodeId;"`

	InstanceTable
}

type InstanceTable struct{}

func (i InstanceTable) BeforeCreate(db *gorm.DB) error {
	return db.Set(tableOptions, fmt.Sprintf("comment '%s'", i.TableComment())).Error
}

func (i InstanceTable) TableName() string {
	return "instances"
}

func (i InstanceTable) TableComment() string {
	return "remote instance record table"
}

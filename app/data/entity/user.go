package entity

import (
	"time"
)

type User struct {
	UUID     string    `gorm:"column:uuid;comment:User UUID;primaryKey;index;type:varchar(256);"`
	Username string    `gorm:"column:username;comment:username;uniqueIndex;type:varchar(256);"`
	Password string    `gorm:"column:password;comment:User password;type:varchar(256);"`
	Email    string    `gorm:"column:email;comment:User concat email;type:varchar(256);"`
	CreateAt time.Time `gorm:"column:create_at;comment:User create time"`
}

func (u User) TableName() string {
	return "users"
}

func (u User) TableComment() string {
	return "system users entity table"
}

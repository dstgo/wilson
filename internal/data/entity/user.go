package entity

import (
	"gorm.io/gorm"
)

// User user entity represents from database table
type User struct {
	UUID     string `gorm:"uniqueIndex;type:varchar(256);comment:User UUID;"`
	Username string `gorm:"uniqueIndex;type:varchar(256);comment:username;"`
	Password string `gorm:"comment:User password;type:varchar(256);"`
	Email    string `gorm:"uniqueIndex;type:varchar(256);comment:User concat email;"`

	gorm.Model
	UserTable
}

type UserTable struct{}

func (u UserTable) TableName() string {
	return "users"
}

func (u UserTable) TableComment() string {
	return "app users entity table"
}

// UserRole user-role relation table
type UserRole struct {
	UserId uint `gorm:"primaryKey;comment:id of user who own role;"`
	User   User `gorm:"foreignKey:UserId;"`

	RoleId uint `gorm:"primaryKey;comment:id of role;"`
	Role   Role `gorm:"foreignKey:RoleId;"`

	UserRoleTable
}

type UserRoleTable struct{}

func (u UserRoleTable) TableName() string {
	return "users_roles"
}

func (u UserRoleTable) TableComment() string {
	return "user-role relation table"
}

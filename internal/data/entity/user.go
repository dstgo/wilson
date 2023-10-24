package entity

import (
	"fmt"
	"gorm.io/gorm"
)

// User user entity represents from database table
type User struct {
	gorm.Model
	UserTable

	UUID     string `gorm:"uniqueIndex;type:varchar(40);comment:User UUID;"`
	Username string `gorm:"uniqueIndex;type:varchar(30);comment:username;"`
	Password string `gorm:"comment:User password;type:varchar(255);"`
	Email    string `gorm:"uniqueIndex;type:varchar(80);comment:User concat email;"`

	Instances []Instance `gorm:"foreignKey:UserId;"`
	Roles     []Role     `gorm:"many2many:users_roles;"`
}

type UserTable struct{}

func (u UserTable) BeforeCreate(db *gorm.DB) error {
	return db.Set(tableOptions, fmt.Sprintf("comment '%s'", u.TableComment())).Error
}

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

func (u UserRole) BeforeCreate(db *gorm.DB) error {
	return db.SetupJoinTable(User{}, "Roles", UserRole{})
}

type UserRoleTable struct{}

func (u UserRoleTable) BeforeCreate(db *gorm.DB) error {
	return db.Set(tableOptions, fmt.Sprintf("comment '%s'", u.TableComment())).Error
}

func (u UserRoleTable) TableName() string {
	return "users_roles"
}

func (u UserRoleTable) TableComment() string {
	return "user-role relation table"
}

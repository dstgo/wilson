package entity

import "gorm.io/gorm"

// Role app roles record table
type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);comment:role display name;"`
	Code string `gorm:"type:varchar(255);comment:role code;"`
	RoleTable
}

type RoleTable struct{}

func (r RoleTable) TableName() string {
	return "roles"
}

func (r RoleTable) TableComment() string {
	return "app roles record table"
}

// Permission app permissions record table
type Permission struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);comment:perm name;"`
	Object string `gorm:"type:varchar(255);comment:resource will be accessed;"`
	Action string `gorm:"type:varchar(255);comment:resource action;"`
	Group  string `gorm:"type:varchar(255);comment:permission group;"`
	Tag    string `gorm:"type:varchar(255);comment:perm's tag,define type of perm;"`

	PermissionTable
}

type PermissionTable struct{}

func (p PermissionTable) TableName() string {
	return "permissions"
}

func (p PermissionTable) TableComment() string {
	return "app permissions record table"
}

// RolePermission role-permission relation table
type RolePermission struct {
	PermId     uint       `gorm:"primaryKey;comment:id of permission;"`
	Permission Permission `gorm:"foreignKey:PermId;"`

	RoleId uint `gorm:"primaryKey;comment:id of role;"`
	Role   Role `gorm:"foreignKey:RoleId;"`

	RolePermissionTable
}

type RolePermissionTable struct{}

func (r RolePermissionTable) TableName() string {
	return "roles_permissions"
}

func (r RolePermissionTable) TableComment() string {
	return "role-permission relation table"
}

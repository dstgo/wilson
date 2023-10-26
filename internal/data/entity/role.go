package entity

import (
	"fmt"
	"gorm.io/gorm"
)

// Role app roles record table
type Role struct {
	RoleTable
	Id   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(50);comment:role display name;"`
	Code string `gorm:"type:varchar(50);uniqueIndex;comment:role code;"`

	CreatedAt uint64 `gorm:"autoCreateTime:nano;"`
	UpdatedAt uint64 `gorm:"autoUpdateTime:nano;"`

	Perms []Permission `gorm:"many2many:role_permission;"`
	Users []User       `gorm:"many2many:users_roles;"`
}

type RoleTable struct{}

func (r RoleTable) BeforeCreate(db *gorm.DB) error {
	return db.Set(tableOptions, fmt.Sprintf("comment '%s'", r.TableComment())).Error
}

func (r RoleTable) TableName() string {
	return "roles"
}

func (r RoleTable) TableComment() string {
	return "app roles record table"
}

// Permission app permissions record table
type Permission struct {
	PermissionTable

	Id     uint   `gorm:"primarykey"`
	Name   string `gorm:"type:varchar(50);comment:perm name;"`
	Object string `gorm:"type:varchar(100);uniqueIndex:perm;comment:resource will be accessed;"`
	Action string `gorm:"type:varchar(50);uniqueIndex:perm;comment:resource action;"`
	Group  string `gorm:"type:varchar(30);uniqueIndex:perm;comment:permission group;"`
	Tag    string `gorm:"type:varchar(30);uniqueIndex:perm;comment:perm's tag,define type of perm;"`

	CreatedAt uint64 `gorm:"autoCreateTime:nano;"`
	UpdatedAt uint64 `gorm:"autoUpdateTime:nano;"`

	Roles []Role `gorm:"many2many:role_permission;"`
}

type PermissionTable struct{}

func (p PermissionTable) BeforeCreate(db *gorm.DB) error {
	return db.Set(tableOptions, fmt.Sprintf("comment '%s'", p.TableComment())).Error
}

func (p PermissionTable) TableName() string {
	return "permissions"
}

func (p PermissionTable) TableComment() string {
	return "app permissions record table"
}

// RolePermission role-permission relation table
type RolePermission struct {
	PermissionId uint       `gorm:"primaryKey;comment:id of permission;"`
	Permission   Permission `gorm:"foreignKey:PermissionId;"`

	RoleId uint `gorm:"primaryKey;comment:id of role;"`
	Role   Role `gorm:"foreignKey:RoleId;"`

	RolePermissionTable
}

func (r RolePermission) BeforeCreate(db *gorm.DB) error {
	db.Set(tableOptions, fmt.Sprintf("comment '%s'", r.TableComment()))
	return db.SetupJoinTable(Role{}, "Perms", RolePermission{})
}

type RolePermissionTable struct{}

func (r RolePermissionTable) TableName() string {
	return "role_permission"
}

func (r RolePermissionTable) TableComment() string {
	return "role-permission relation table"
}

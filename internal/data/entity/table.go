package entity

import (
	"gorm.io/gorm"
)

const (
	tableOptions = "gorm:table_options"
)

// Table
// represents an entity for table
type Table interface {
	TableName() string
	TableComment() string
}

var tables = []any{
	// user
	&User{},
	// node
	&Node{},
	// instance
	&Instance{},
	// role
	&Role{},
	&Permission{},

	// relation table
	&RolePermission{},
	&UserRole{},
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(tables...)
}

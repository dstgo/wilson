package entity

import (
	"fmt"
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

var tables = []Table{
	// user
	User{},
	UserRole{},
	// node
	Node{},
	// instance
	Instance{},
	// role
	Role{},
	Permission{},
	RolePermission{},
}

func Migrate(db *gorm.DB) error {
	for _, table := range tables {
		err := db.Set(tableOptions, fmt.Sprintf("comment '%s'", table.TableComment())).
			Migrator().
			AutoMigrate(table)
		if err != nil {
			return err
		}
	}
	return nil
}

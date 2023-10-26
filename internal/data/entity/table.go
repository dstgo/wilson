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

// Migrate auto migration all the table defined
func Migrate(db *gorm.DB) error {
	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			err := db.Set(tableOptions, fmt.Sprintf(`comment '%s'`, table.TableComment())).Migrator().CreateTable(table)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

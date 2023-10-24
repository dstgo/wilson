package data

import (
	"github.com/dstgo/wilson/internal/data/entity"
	"testing"
)

func TestMigrate(t *testing.T) {
	db, err := DialTestDB()
	if err != nil {
		t.Error(err)
		return
	}
	db.AutoMigrate(entity.Role{}, entity.Permission{}, entity.RolePermission{})

	roles := []entity.Role{{Name: "admin", Code: "admin"}}

	perms := []entity.Permission{{Name: "a", Object: "b", Action: "c"}}

	db.Create(&roles)
	db.Create(&perms)

	db.Model(&roles).Association("Perms").Append(&perms)
}

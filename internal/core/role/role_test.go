package role

import (
	"github.com/dstgo/wilson/internal/data/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func connect(t *testing.T) *gorm.DB {
	dsn := "root:123456@tcp(192.168.48.138:3306)/wilson?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		t.Error("db connect error", err)
	}
	t.Log("db connect success")
	return db
}

func TestPermUpsert(t *testing.T) {
	db := connect(t)
	if err := entity.Migrate(db); err != nil {
		t.Error(err)
	}
	perms := []entity.Permission{
		{Object: "/user/info", Action: "GET", Group: "user"},
		{Object: "/user/create", Action: "POST", Group: "user"},
		{Object: "/user/update", Action: "POST", Group: "user"},
	}

	batch, err := createPermInBatch(db, perms)
	if err != nil {
		t.Error(err)
	}
	t.Log(batch)
}

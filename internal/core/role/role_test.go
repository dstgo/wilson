package role

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"testing"
)

func TestPermUpsert(t *testing.T) {
	db, err := data.DialTestDB()
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

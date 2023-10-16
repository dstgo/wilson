package data

import (
	"fmt"
	"gorm.io/gorm"
)

func Page(db *gorm.DB, page, size int) *gorm.DB {
	if page <= 0 || size < 0 {
		panic(fmt.Sprintf("invliad page parameter: page %d, size %d", page, size))
	}
	return db.Offset((page - 1) * size).Limit(size)
}

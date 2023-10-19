package data

import (
	"fmt"
	"gorm.io/gorm"
)

type GormOption func(*gorm.DB) *gorm.DB

func Pages(page, size int) GormOption {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 || size < 0 {
			panic(fmt.Sprintf("invliad page parameter: page %d, size %d", page, size))
		}
		return db.Offset((page - 1) * size).Limit(size)
	}
}

func Order(column string, desc bool) GormOption {
	return func(db *gorm.DB) *gorm.DB {
		var descStr string
		if desc {
			descStr = "DESC"
		}
		return db.Order(fmt.Sprintf("%s %s", column, descStr))
	}
}

const LikeDelim = "%"

func Like(pattern string) string {
	return fmt.Sprintf("%s%s%s", LikeDelim, pattern, LikeDelim)
}

func LikeSuffix(pattern string) string {
	return fmt.Sprintf("%s%s", pattern, LikeDelim)
}

func LikePrefix(pattern string) string {
	return fmt.Sprintf("%s%s", LikeDelim, pattern)
}

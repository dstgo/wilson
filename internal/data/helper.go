package data

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DialTestDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(192.168.48.138:3306)/wilson?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return db, err
	}
	return db, nil
}

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

func HasRecordFound(db *gorm.DB) (bool, error) {
	err := db.Error
	rows := db.RowsAffected

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	} else {
		return rows > 0, nil
	}
}

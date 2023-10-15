package data

import "gorm.io/gorm"

func Paginator(page, size int, conds string, args ...any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page-1)*size).Limit(size).Where(conds, args...)
	}
}

func PaginatorWithOrder(page, size int, order any, conds string, args ...any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order).Offset((page-1)*size).Limit(size).Where(conds, args...)
	}
}

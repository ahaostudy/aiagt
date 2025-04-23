package gormutil

import "gorm.io/gorm"

type Scope func(db *gorm.DB) *gorm.DB

func Where(query interface{}, args ...interface{}) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

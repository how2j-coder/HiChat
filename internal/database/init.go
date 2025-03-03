package database

import (
	"gorm.io/gorm"
	"sync"
)

var (
	db *gorm.DB
	dbOnce sync.Once
)

// InitDB connect database
func InitDB()  {
	db = InitMysql()
}

func GetDB() *gorm.DB {
	if db == nil {
		dbOnce.Do(func() {
			InitDB()
		})
	}

	return db
}

// CloseMysql close db
func CloseMysql() error {
	return DbClose(db)
}

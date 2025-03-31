package database

import (
	"com/chat/service/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// InitDB connect database
func InitDB() {
	db = InitMysql()
}

// InitTables 初始创建表
func InitTables() error {
	err := db.Session(
		&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}, // 关闭自动创建表时的日志打印
	).AutoMigrate(
		model.User{},
		model.Role{},
		model.RoleUser{},
		model.Menu{},
		model.RoleMenu{},
		model.CasbinRule{},
		model.Platform{},
		model.Upload{},
	)
	return err
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

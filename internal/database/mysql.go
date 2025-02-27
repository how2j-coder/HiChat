package database

import (
	"com/chat/service/internal/config"
	"database/sql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql()  (*gorm.DB, error ) {
	var dsn string
	mysqlCfg := config.GetConfig().Database.Mysql
	dsn = mysqlCfg.Dsn
	sqlDB, _ := sql.Open("mysql", dsn)
	db, err := gorm.Open(
		mysqlDriver.New(mysqlDriver.Config{
			Conn: sqlDB,
		}),
		)
	if err != nil {
		return nil, err
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	return db, nil
}

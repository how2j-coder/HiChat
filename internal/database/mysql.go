package database

import (
	"com/chat/service/internal/config"
	"database/sql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql()  *gorm.DB {
	var dsn string
	mysqlCfg := config.GetConfig().Database.Mysql
	dsn = mysqlCfg.Dsn
	sqlDB, _ := sql.Open("mysql", dsn)
	mdb, err := gorm.Open(
		mysqlDriver.New(mysqlDriver.Config{
			Conn: sqlDB,
		}),
		)
	if err != nil {
		panic("init mysql error: " + err.Error())
		return nil
	}
	mdb.Set("gorm:table_options", "CHARSET=utf8mb4")
	return mdb
}
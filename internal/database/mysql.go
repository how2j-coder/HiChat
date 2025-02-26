package database

import (
	"database/sql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(dsn string)  (*gorm.DB, error ) {
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

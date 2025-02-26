package database

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func InitDB()  {
	db = InitMysql()
}
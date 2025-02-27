package database

import "gorm.io/gorm"

var (
	db *gorm.DB
)

// InitDB connect database
func InitDB()  {
	_, _ = InitMysql()
}
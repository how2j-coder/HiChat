package initialize

import (
	"HiChat/global"
	"fmt"
)

func InitDB() {
	dbConfig := global.ServiceConfig.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	global.Logger.Sugar().Info(dsn)
}

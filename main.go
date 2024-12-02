package main

import (
	"HiChat/global"
	"HiChat/initialize"
)

func main() {
	initialize.InitConfig() // 配置
	initialize.InitLogger() // 日志
	initialize.InitDB()     // 数据库
	initialize.InitRedis()  // redis 数据库
	initialize.InitRouter() // 路由
	global.Logger.Sugar().Info("server running")
}

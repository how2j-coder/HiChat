package main

import (
	"HiChat/global"
	"HiChat/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()

	global.Logger.Info("how2j")
}

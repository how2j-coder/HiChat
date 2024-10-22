package main

import (
	"HiChat/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
}

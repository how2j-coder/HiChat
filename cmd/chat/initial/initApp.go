package initial

import (
	"com/chat/service/configs"
	"com/chat/service/internal/config"
	"com/chat/service/internal/database"
	"flag"
	"fmt"
)

var (
	version string
	configFile string
)

func InitApp() {
	initConfig()
	database.InitDB()
	database.InitCache()
}

// 初始化配置
func initConfig()  {
	flag.StringVar(&version, "version", "", "service Version Number")
	flag.StringVar(&configFile, "c", "", "configuration file")
	flag.Parse()

	getConfigFormLocal()

	fmt.Println(version, "version")
}

func getConfigFormLocal()  {
	if configFile == "" {
		configFile = configs.Path("server.yml")
	}
	err := config.Init(configFile)
	if err != nil {
		panic("init config error: " + err.Error())
	}
}
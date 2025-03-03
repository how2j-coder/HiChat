package initial

import (
	"com/chat/service/configs"
	"com/chat/service/internal/config"
	"com/chat/service/internal/database"
	"com/chat/service/pkg/logger"
	"flag"
)

var (
	version    string
	configFile string
)

func InitApp() {
	initConfig()

	cfg := config.GetConfig()
	_, err := logger.Init(
		logger.WithLevel(cfg.Logger.Level),
		logger.WithFormat(cfg.Logger.Format),
	)

	if err != nil {
		panic(err)
	}

	logger.Debug(config.Show())
	logger.Info("[logger] was initialized")

	database.InitDB()
	logger.InfoPf("[%s] was initialized", cfg.Database.Driver)

	database.InitCache()
	logger.Info("[redis] was initialized")
}

// 初始化配置
func initConfig() {
	flag.StringVar(&version, "version", "", "service Version Number")
	flag.StringVar(&configFile, "c", "", "configuration file")
	flag.Parse()

	getConfigFormLocal()
}

func getConfigFormLocal() {
	if configFile == "" {
		// 配置文件
		configFile = configs.Path("server.yml")
	}
	err := config.Init(configFile)
	if err != nil {
		panic("init config error: " + err.Error())
	}
}
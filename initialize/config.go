package initialize

import (
	"HiChat/global"
	"HiChat/utils"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	// config 文件路径
	rootPath := utils.GetRootPath()
	configFile := fmt.Sprintf("%s/config_dev.yaml", rootPath)

	// viper 实例化
	v := viper.New()
	// set config file
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)

	// read config file
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// put data into global.ServiceConfig, global use
	if err := v.Unmarshal(&global.ServiceConfig); err != nil {
		panic(err)
	}
	zap.S().Info("配置信息", global.ServiceConfig)

}

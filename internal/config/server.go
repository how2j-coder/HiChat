package config

import "com/chat/service/pkg/conf"

var config *Config

func Init(configFile string, fs ...func()) error  {
	config = &Config{}
	return conf.Parse(configFile, config, fs...)
}

func GetConfig() *Config  {
	if config == nil {
		panic("config is nil, please call config.Init() first")
	}
	return config
}

type App struct {
	Host string `yaml:"host" json:"host"`
	Env string `yaml:"env" json:"env"`
	Name string `yaml:"name" json:"name"`
}

type Config struct {
	App   App `yaml:"app" json:"app"`
}
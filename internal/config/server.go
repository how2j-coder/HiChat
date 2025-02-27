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

type Mysql struct {
	Dsn             string   `yaml:"dsn" json:"dsn"`
	EnableLog       bool     `yaml:"enableLog" json:"enableLog"`
}

type Redis struct {
	DialTimeout  int    `yaml:"dialTimeout" json:"dialTimeout"`
	Dsn          string `yaml:"dsn" json:"dsn"`
	ReadTimeout  int    `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout" json:"writeTimeout"`
}

type Database struct {
	Driver     string  `yaml:"driver" json:"driver"`
	Mysql      Mysql   `yaml:"mysql" json:"mysql"`
}

type Config struct {
	App   App `yaml:"app" json:"app"`
	Database Database `yaml:"database" json:"database"`
	Redis 	Redis   `yaml:"redis" json:"redis"`
}
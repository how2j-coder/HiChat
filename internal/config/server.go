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

func Show(hiddenFields ...string) string {
	return conf.Show(config, hiddenFields...)
}

type Config struct {
	App   App `yaml:"app" json:"app"`
	Database Database `yaml:"database" json:"database"`
	Redis 	Redis   `yaml:"redis" json:"redis"`
	Logger Logger `yaml:"logger" json:"logger"`
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
type Logger struct {
	Level string `yaml:"level" json:"level"`
	Format string `yaml:"format" json:"format"`
	IsSave bool   `yaml:"isSave" json:"isSave"`
	LogFileConfig LogFileConfig `yaml:"logFileConfig" json:"logFileConfig"`
}

type LogFileConfig struct {
	Filename   string `yaml:"filename" json:"filename"`
	MaxSize    int    `yaml:"maxSize" json:"maxSize"`
	MaxBackups int    `yaml:"maxBackups" json:"maxBackups"`
	MaxAge     int    `yaml:"maxAge" json:"maxAge"`
	IsCompression bool   `yaml:"isCompression" json:"isCompression"`
}


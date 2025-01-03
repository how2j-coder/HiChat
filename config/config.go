package config

type ServiceConfig struct {
	DB   dbSqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisDB RedisConfig `mapstructure:"redis" json:"redis"`
	Port int         `mapstructure:"port" json:"port"`
	Log  Log         `mapstructure:"log" json:"log"`
}

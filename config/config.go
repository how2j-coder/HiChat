package config

type ServiceConfig struct {
	DB   dbSqlConfig `mapstructure:"mysql" json:"mysql"`
	Port int         `mapstructure:"port" json:"port"`
}

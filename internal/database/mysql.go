package database

import (
	"com/chat/service/internal/config"
	"com/chat/service/pkg/datastore/mysql"
	"com/chat/service/pkg/logger"
	"com/chat/service/pkg/utils"
	"gorm.io/gorm"
	"time"
)

func InitMysql() *gorm.DB {
	mysqlCfg := config.GetConfig().Database.Mysql
	opts := []mysql.Option{
		mysql.WithMaxIdleConns(mysqlCfg.MaxIdleConns),
		mysql.WithMaxOpenConns(mysqlCfg.MaxOpenConns),
		mysql.WithConnMaxLifetime(time.Duration(mysqlCfg.ConnMaxLifetime) * time.Minute),
	}

	if mysqlCfg.EnableLog {
		opts = append(opts,
			mysql.WithLogging(logger.Get()),
			mysql.WithLogRequestIDKey("request_id"),
		)
	}

	dsn := utils.AdaptiveMysqlDsn(mysqlCfg.Dsn)
	dbI, err := mysql.Init(dsn, opts...)
	if err != nil {
		panic("init mysql error: " + err.Error())
	}
	return dbI
}

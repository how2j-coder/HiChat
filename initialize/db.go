package initialize

import (
	"HiChat/global"
	"HiChat/models"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func CreateDBTable(db *gorm.DB, tab ...interface{}) error {
	for _, v := range tab {
		exist := db.Migrator().HasTable(v)
		if !exist {
			err := db.Migrator().CreateTable(v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func InitDB() {
	dbConfig := global.ServiceConfig.DB
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name,
	)

	//写sql语句配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	var err error
	//将获取到的连接赋值到global.DB
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, //打印sql日志
	})
	if err != nil {
		panic(err)
	} else {
		err := CreateDBTable(global.DB,
			&models.User{}, &models.File{},
			&models.Platform{}, &models.Menu{},
		)
		if err != nil {
			global.Logger.Sugar().Error("Failed to connect to Mysql",err.Error())
			return
		} else {
			global.Logger.Info("mysql successfully initialized")
		}
	}
}

func InitRedis()  {
	opt := redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			global.ServiceConfig.RedisDB.Host,
			global.ServiceConfig.RedisDB.Port,
			),
			Password: global.ServiceConfig.RedisDB.Password,
			DB: 0,
	}
	rdb := redis.NewClient(&opt)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		global.Logger.Sugar().Error("Failed to connect to Redis:",err.Error())
		panic(err)
		return
	}else {
		global.Logger.Info("redis successfully initialized")
	}
	global.RedisDB = rdb
}
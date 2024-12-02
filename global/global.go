package global

import (
	"HiChat/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ServiceConfig *config.ServiceConfig
	DB            *gorm.DB
	RedisDB       *redis.Client
	Logger        *zap.Logger
)

const AuthCtxFiled = "userId"
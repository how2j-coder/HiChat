package global

import (
	"HiChat/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ServiceConfig *config.ServiceConfig
	DB            *gorm.DB
	Logger        *zap.Logger
)

const AuthCtxFiled = "userId"
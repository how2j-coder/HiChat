package database

import (
	"com/chat/service/internal/config"
	storeRedis "com/chat/service/pkg/datastore/redis"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	redisClient  *redis.Client
	redisCliOnce sync.Once
)

func InitCache() {
	GetRedisClient()
}

func InitRedis() {
	redisCfg := config.GetConfig().Redis
	opts := []storeRedis.Option{
		storeRedis.WithDialTimeout(time.Duration(redisCfg.DialTimeout) * time.Second),
		storeRedis.WithReadTimeout(time.Duration(redisCfg.ReadTimeout) * time.Second),
		storeRedis.WithWriteTimeout(time.Duration(redisCfg.WriteTimeout) * time.Second),
	}

	var err error
	redisClient, err = storeRedis.Init(redisCfg.Dsn, opts...)
	if err != nil {
		panic("redis.Init error: " + err.Error())
	}
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		redisCliOnce.Do(func() {
			InitRedis()
		})
	}
	return redisClient
}

// CloseRedis close redis
func CloseRedis() error {
	return DbClose(redisClient)
}

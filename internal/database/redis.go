package database

import (
	"com/chat/service/internal/config"
	"context"
	"github.com/redis/go-redis/v9"
	"strings"
	"sync"
	"time"
)

var (
	redisClient *redis.Client
	redisCliOnce sync.Once
)

func InitCache() {
	InitRedis()
}

func InitRedis() {
	redisConfig := config.GetConfig().Redis
	dsn := redisConfig.Dsn
	dsn = strings.ReplaceAll(dsn, " ", "")
	if len(dsn) > 8 {
		if !strings.Contains(dsn[len(dsn)-3:], "/") {
			dsn += "/0" // use db 0 by default
		}

		if dsn[:8] != "redis://" && dsn[:9] != "redis://" {
			dsn = "redis://" + dsn
		}
	}

	opts, err := redis.ParseURL(dsn)

	if err != nil {
		panic("redis.Init error: " + err.Error())
	}

	if redisConfig.DialTimeout > 0 {
		opts.DialTimeout = time.Duration(redisConfig.DialTimeout) * time.Second
	}
	if redisConfig.ReadTimeout > 0 {
		opts.ReadTimeout = time.Duration(redisConfig.ReadTimeout) * time.Second
	}
	if redisConfig.WriteTimeout > 0 {
		opts.WriteTimeout = time.Duration(redisConfig.WriteTimeout) * time.Second
	}

	rdb := redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	err = rdb.Ping(ctx).Err()

	redisClient = rdb

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

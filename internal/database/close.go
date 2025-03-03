package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

// DbClose close mysql & mysql
func DbClose(db interface{}) error {
	switch db := db.(type) {
	case *gorm.DB:
		return MysqlClose(db)
	case *redis.Client:
		return RedisClose(db)
		// 其他数据库类型的关闭方法
	default:
		return fmt.Errorf("unsupported database type: %T", db)
	}
}

// RedisClose Redis client
func RedisClose(rdb *redis.Client) error {
	if rdb == nil {
		return nil
	}

	err := rdb.Close()
	if err != nil && errors.Is(err, redis.ErrClosed) {
		return err
	}

	return nil
}

// MysqlClose gorm db
func MysqlClose(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	checkInUse(sqlDB, time.Second*5)

	return sqlDB.Close()
}

func checkInUse(sqlDB *sql.DB, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	for {
		select {
		case <-time.After(time.Millisecond * 250):
			if v := sqlDB.Stats().InUse; v == 0 {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}


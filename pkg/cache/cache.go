package cache

import (
	"context"
	"errors"
	"time"
)

var (
	// DefaultExpireTime default expiry time 默认到期时间
	DefaultExpireTime = time.Hour * 24
	// DefaultNotFoundExpireTime expiry time when result is empty 1 minute, result为空时的过期时间 1 分钟,
	// often used for cache time when data is empty (cache pass-through) 通常用于数据为空时的缓存时间（缓存直通）.
	DefaultNotFoundExpireTime = time.Minute * 10

	// NotFoundPlaceholder placeholder
	NotFoundPlaceholder      = "*"
	NotFoundPlaceholderBytes = []byte(NotFoundPlaceholder)
	ErrPlaceholder           = errors.New("cache: placeholder")

	// DefaultClient generate a cache client, where keyPrefix is generally the business prefix 生成缓存客户端, 其中 keyPrefix 一般为业务前缀.
	DefaultClient Cache
)

// Cache driver interfac
// Cache driver interface
type Cache interface {
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, val interface{}) error
	MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error
	MultiGet(ctx context.Context, keys []string, valueMap interface{}) error
	Del(ctx context.Context, keys ...string) error
	SetCacheWithNotFound(ctx context.Context, key string) error
}

// Set data
func Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	return DefaultClient.Set(ctx, key, val, expiration)
}

// Get data
func Get(ctx context.Context, key string, val interface{}) error {
	return DefaultClient.Get(ctx, key, val)
}

// MultiSet multiple set data
func MultiSet(ctx context.Context, valMap map[string]interface{}, expiration time.Duration) error {
	return DefaultClient.MultiSet(ctx, valMap, expiration)
}

// MultiGet multiple get data
func MultiGet(ctx context.Context, keys []string, valueMap interface{}) error {
	return DefaultClient.MultiGet(ctx, keys, valueMap)
}

// Del multiple delete data
func Del(ctx context.Context, keys ...string) error {
	return DefaultClient.Del(ctx, keys...)
}

// SetCacheWithNotFound .
func SetCacheWithNotFound(ctx context.Context, key string) error {
	return DefaultClient.SetCacheWithNotFound(ctx, key)
}

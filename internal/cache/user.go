package cache

import (
	"com/chat/service/pkg/cache"
	"com/chat/service/pkg/datastore/redis"
	"com/chat/service/pkg/encoding"
	"com/chat/service/pkg/utils"
	"context"
	"errors"
	"time"
)

const (
	userCachePrefixKey  = "userCache:"
	UserCacheExpireTime = 5 * time.Minute // 5分钟
)

type UserCache interface {
	Set(ctx context.Context, id uint64,data *string, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*string, error)
	MustGet(ctx context.Context, ids []uint64) (map[uint64]*string, error)
	MustSet(ctx context.Context, ids map[uint64]*string, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// UserCache define a cache struct
type userCache struct {
	cache cache.Cache
}

var _ UserCache = (*userCache)(nil)

// NewUserCache new a cache
func NewUserCache(cacheRDB *redis.Client) UserCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	c := cache.NewRedisCache(cacheRDB, cachePrefix, jsonEncoding, func() interface{} {
		return ""
	})
	return &userCache{cache: c}
}

func (c *userCache) Set(ctx context.Context, id uint64, data *string, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

func (c *userCache) Get(ctx context.Context, id uint64) (*string, error) {
	var data *string
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *userCache) MustGet(ctx context.Context, ids []uint64) (map[uint64]*string, error) {
	var keys []string
	for _, id := range ids {
		cacheKey := c.GetUserCacheKey(id)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*string, len(keys))
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*string, len(keys))
	for _, id := range ids {
		val, ok := itemMap[c.GetUserCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}
	return retMap, nil
}

func (c *userCache) MustSet(ctx context.Context, ids map[uint64]*string, duration time.Duration) error {
	valMap := make(map[string]interface{}, len(ids))
	for key, val := range ids {
		cacheKey := c.GetUserCacheKey(key)
		valMap[cacheKey] = val
	}
	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}
	return nil
}

func (c *userCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

func (c *userCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

func (c *userCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}


// GetUserCacheKey cache key
func (c *userCache) GetUserCacheKey(id uint64) string {
	return userCachePrefixKey + utils.Uint64ToStr(id)
}
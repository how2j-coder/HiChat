package cache

import (
	"time"
)

const (
	userCachePrefixKey  = "userCache:"
	UserCacheExpireTime = 5 * time.Minute // 5分钟
)

// userExampleCache define a cache struct
type userExampleCache struct {
}

type UserExampleCache interface {
}

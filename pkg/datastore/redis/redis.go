package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

// Client is a redis client
type Client = redis.Client

// Init connecting to redis
// dsn supported formats.
// (1) no password, no db: localhost:6379
// (2) with a password and db: <user>:<pass>@localhost:6379/2
// (3) redis://default:123456@localhost:6379/0?max_retries=3
// for more parameters see the redis source code for the setupConnParams function.
func Init(dsn string, opts ...Option) (*redis.Client, error) {
	o := defaultOptions()
	o.apply(opts...)

	opt, err := getRedisOpt(dsn, o)
	if err != nil {
		return nil, err
	}

	// replace single options if provided
	if o.singleOptions != nil {
		opt = o.singleOptions
	}

	rdb := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // nolint
	defer cancel()
	err = rdb.Ping(ctx).Err()

	return rdb, err
}

// InitSingle connecting to single redis instance
func InitSingle(addr string, password string, db int, opts ...Option) (*redis.Client, error) {
	o := defaultOptions()
	o.apply(opts...)

	opt := &redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		DialTimeout:  o.dialTimeout,
		ReadTimeout:  o.readTimeout,
		WriteTimeout: o.writeTimeout,
		TLSConfig:    o.tlsConfig,
	}

	// replace single options if provided
	if o.singleOptions != nil {
		opt = o.singleOptions
	}

	rdb := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // nolint
	defer cancel()
	err := rdb.Ping(ctx).Err()

	return rdb, err
}

func getRedisOpt(dsn string, opts *options) (*redis.Options, error) {
	dsn = strings.ReplaceAll(dsn, " ", "")
	if len(dsn) > 8 {
		if !strings.Contains(dsn[len(dsn)-3:], "/") {
			dsn += "/0" // use db 0 by default
		}

		if dsn[:8] != "redis://" && dsn[:9] != "redis://" {
			dsn = "redis://" + dsn
		}
	}

	redisOpts, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	if opts.dialTimeout > 0 {
		redisOpts.DialTimeout = opts.dialTimeout
	}
	if opts.readTimeout > 0 {
		redisOpts.ReadTimeout = opts.readTimeout
	}
	if opts.writeTimeout > 0 {
		redisOpts.WriteTimeout = opts.writeTimeout
	}
	if opts.tlsConfig != nil {
		redisOpts.TLSConfig = opts.tlsConfig
	}

	return redisOpts, nil
}

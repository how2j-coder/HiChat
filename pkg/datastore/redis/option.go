package redis

import (
	"crypto/tls"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/sdk/trace"
	"time"
)

type options struct {
	dialTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
	tlsConfig    *tls.Config

	// Note: this field is only used for Init and InitSingle, and the other parameters will be ignored.
	singleOptions *redis.Options

	// Note: this field is only used for InitSentinel, and the other parameters will be ignored.
	sentinelOptions *redis.FailoverOptions

	// Note: this field is only used for InitCluster, and the other parameters will be ignored.
	clusterOptions *redis.ClusterOptions
}

// Option set the redis options.
type Option func(*options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// default settings
func defaultOptions() *options {
	return &options{}
}

// WithDialTimeout set dail timeout
func WithDialTimeout(t time.Duration) Option {
	return func(o *options) {
		o.dialTimeout = t
	}
}

// WithReadTimeout set read timeout
func WithReadTimeout(t time.Duration) Option {
	return func(o *options) {
		o.readTimeout = t
	}
}

// WithWriteTimeout set write timeout
func WithWriteTimeout(t time.Duration) Option {
	return func(o *options) {
		o.writeTimeout = t
	}
}

// WithTLSConfig set TLS config
func WithTLSConfig(c *tls.Config) Option {
	return func(o *options) {
		o.tlsConfig = c
	}
}

// WithSingleOptions set single redis options
func WithSingleOptions(opt *redis.Options) Option {
	return func(o *options) {
		o.singleOptions = opt
	}
}

// WithSentinelOptions set redis sentinel options
func WithSentinelOptions(opt *redis.FailoverOptions) Option {
	return func(o *options) {
		o.sentinelOptions = opt
	}
}

// WithClusterOptions set redis cluster options
func WithClusterOptions(opt *redis.ClusterOptions) Option {
	return func(o *options) {
		o.clusterOptions = opt
	}
}

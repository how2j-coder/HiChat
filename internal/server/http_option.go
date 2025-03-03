package server

import "com/chat/service/pkg/servicerd/registry"

// HTTPOption setting up http
type HTTPOption func(*httpOptions)

type httpOptions struct {
	isProd    bool  // 是否为开发环境
	instance  *registry.ServiceInstance  // 服务实例
	iRegistry registry.Registry // 是否注册
}

// 服务的基本数据
func defaultHTTPOptions() *httpOptions {
	return &httpOptions{
		isProd:    false,
		instance:  nil,
		iRegistry: nil,
	}
}

func (o *httpOptions) apply(opts ...HTTPOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithHTTPIsProd setting up production environment markers
func WithHTTPIsProd(isProd bool) HTTPOption {
	return func(o *httpOptions) {
		o.isProd = isProd
	}
}

// WithHTTPRegistry registration services
func WithHTTPRegistry(iRegistry registry.Registry, instance *registry.ServiceInstance) HTTPOption {
	return func(o *httpOptions) {
		o.iRegistry = iRegistry
		o.instance = instance
	}
}

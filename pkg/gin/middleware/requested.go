package middleware

import (
	"com/chat/service/pkg/srand"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	// ContextRequestIDKey request id for context
	ContextRequestIDKey = "request_id"
	// HeaderXRequestIDKey header request id key
	HeaderXRequestIDKey = "X-Request-Id"
)

// RequestHeaderKey request header key
var RequestHeaderKey = "request_header_key"

// WrapCtx wrap context, put the Keys and Header of gin.Context into context.
func WrapCtx(c *gin.Context) context.Context {
	ctx := context.WithValue(
		c.Request.Context(), ContextRequestIDKey,
		c.GetString(ContextRequestIDKey),
	)
	return context.WithValue(ctx, RequestHeaderKey, c.Request.Header)
}

// ----------------------------------
type requestIDOptions struct {
	contextRequestIDKey string
	headerXRequestIDKey string
}
// RequestIDOption set the request id  options.
type RequestIDOption func(*requestIDOptions)

func defaultRequestIDOptions() *requestIDOptions {
	return &requestIDOptions{
		contextRequestIDKey: ContextRequestIDKey,
		headerXRequestIDKey: HeaderXRequestIDKey,
	}
}
func (o *requestIDOptions) apply(opts ...RequestIDOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *requestIDOptions) setRequestIDKey() {
	if o.contextRequestIDKey != ContextRequestIDKey {
		ContextRequestIDKey = o.contextRequestIDKey
	}
	if o.headerXRequestIDKey != HeaderXRequestIDKey {
		HeaderXRequestIDKey = o.headerXRequestIDKey
	}
}
// RequestID is an interceptor that injects a 'request id' into the context and request/response header of each request.
func RequestID(opts ...RequestIDOption) gin.HandlerFunc {
	// customized request id key
	o := defaultRequestIDOptions()
	o.apply(opts...)
	o.setRequestIDKey()

	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get(HeaderXRequestIDKey)

		// Create request id
		if requestID == "" {
			requestID = srand.String(srand.RAll, 10)
			c.Request.Header.Set(HeaderXRequestIDKey, requestID)
		}

		// Expose it for use in the application
		c.Set(ContextRequestIDKey, requestID)

		// Set X-Request-Id header
		c.Writer.Header().Set(HeaderXRequestIDKey, requestID)

		c.Next()
	}
}

// GCtxRequestID get request id from gin.Context
func GCtxRequestID(c *gin.Context) string {
	if v, isExist := c.Get(ContextRequestIDKey); isExist {
		if requestID, ok := v.(string); ok {
			return requestID
		}
	}
	return ""
}
// GCtxRequestIDField get request id field from gin.Context
func GCtxRequestIDField(c *gin.Context) zap.Field {
	return zap.String(ContextRequestIDKey, GCtxRequestID(c))
}

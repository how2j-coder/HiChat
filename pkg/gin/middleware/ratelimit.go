package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Timeout request time out
func Timeout(d time.Duration) gin.HandlerFunc {
	if d < time.Millisecond {
		return func(c *gin.Context) {}
	}
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), d)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			if c.Writer.Status() == 200 {
				c.AbortWithStatus(http.StatusGatewayTimeout)
				return
			}
			c.Abort()
		}
	}
}
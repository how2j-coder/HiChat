package routers

import (
	"com/chat/service/pkg/gin/middleware"
	"github.com/gin-gonic/gin"
)

var apiV1RouterFns []func(r *gin.RouterGroup)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// register routers, middleware support
	registerRouters(r, "/api/v1", apiV1RouterFns)

	return r
}

func registerRouters(
	r *gin.Engine, groupPath string,
	routerFns []func(*gin.RouterGroup),
	handlers ...gin.HandlerFunc,
) {
	rg := r.Group(groupPath, handlers...)
	for _, fn := range routerFns {
		fn(rg)
	}
}
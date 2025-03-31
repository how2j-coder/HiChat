package routers

import (
	"com/chat/service/internal/config"
	"com/chat/service/pkg/gin/middleware"
	"com/chat/service/pkg/gin/validator"
	"com/chat/service/pkg/jwt"
	"com/chat/service/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

var apiV1RouterFns []func(r *gin.RouterGroup)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	if config.GetConfig().HTTP.Timeout > 0 {
		// if you need more fine-grained control over your routes, set the timeout in your routes, unsetting the timeout globally here.
		r.Use(middleware.Timeout(time.Second * time.Duration(config.GetConfig().HTTP.Timeout)))
	}

	// request id middleware
	r.Use(middleware.RequestID())

	// logger middleware, to print simple messages, replace middleware.Logging with middleware.SimpleLog
	r.Use(middleware.Logging(
		middleware.WithLog(logger.Get()),
		middleware.WithRequestIDFromContext(),
		middleware.WithIgnoreRoutes("/metrics"), // ignore path
	))

	// init jwt middleware
	jwt.Init(
		jwt.WithExpire(time.Hour * 24),
		jwt.WithSigningMethod(jwt.HS384),
	)

	// trace middleware
	// if config.GetConfig().App.EnableTrace {
	// 	r.Use(middleware.Tracing(config.GetConfig().App.Name))
	// }

	// validator
	binding.Validator = validator.Init()

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
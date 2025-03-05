package routers

import (
	"com/chat/service/internal/handler"
	"com/chat/service/pkg/gin/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(rGroup *gin.RouterGroup) {
		userRouter(rGroup, handler.NewUserHandler())
	})
}

func userRouter(group *gin.RouterGroup, h handler.UserHandler) {
	g := group.Group("/user")
	g.Use(middleware.Auth(
		middleware.AddWhiteRouter(map[string]string{
			"/api/v1/user/create": "POST",
			"/api/v1/user/login": "POST",
		}),
	))

	g.POST("/create", h.Create)
	g.POST("/login", h.Login)
	g.PUT("/update/:id", h.UpdateByID)
}

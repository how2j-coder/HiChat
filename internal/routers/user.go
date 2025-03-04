package routers

import (
	"com/chat/service/internal/handler"
	"github.com/gin-gonic/gin"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(rGroup *gin.RouterGroup) {
		userRouter(rGroup, handler.NewUserHandler())
	})
}

func userRouter(group *gin.RouterGroup, h handler.UserHandler) {
	group.POST("/user", h.Create)
}

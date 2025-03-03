package routers

import (
	"com/chat/service/internal/handler"
	"github.com/gin-gonic/gin"
)

func init()  {
	apiV1RouterFns = append(apiV1RouterFns, func(rGroup *gin.RouterGroup) {
		userRouter(rGroup, handler.NewUserHandler())
	})
}

func userRouter(group *gin.RouterGroup, h handler.UserHandler)  {
		group.GET("/login", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"code": 200,
				"msg": "success",
				"data": nil,
			})
		})
}


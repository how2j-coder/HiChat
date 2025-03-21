package routers

import (
	"com/chat/service/internal/handler"
	"github.com/gin-gonic/gin"
)

func init()  {
	apiV1RouterFns = append(apiV1RouterFns, func(rGroup *gin.RouterGroup) {
		platformRouter(rGroup, handler.NewPlatformHandler())
	})
}

func platformRouter(group *gin.RouterGroup, h handler.PlatformHandler) {
	g := group.Group("/platform")
	g.POST("/create", h.Create)
	g.PUT("/update/:id", h.UpdateById)
	g.GET("/list", h.GetColumn)
	g.DELETE("/delete/:id", h.DeleteById)
}
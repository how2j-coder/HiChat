package routers

import (
	"com/chat/service/internal/handler"
	"github.com/gin-gonic/gin"
)

func init()  {
	apiV1RouterFns = append(apiV1RouterFns, func(rGroup *gin.RouterGroup) {
		menuRouter(rGroup, handler.NewMenuHandler())
	})
}

func menuRouter(group *gin.RouterGroup, h handler.MenuHandler)  {
	g := group.Group("/menu")
	g.POST("/create", h.Create)
	g.PUT("/update/:id", h.UpdateByID)
	g.DELETE("/delete/:id", h.DeleteByID)
}
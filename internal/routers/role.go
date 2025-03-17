package routers

import (
	"com/chat/service/internal/handler"
	"github.com/gin-gonic/gin"
)

func init()  {
	apiV1RouterFns = append(apiV1RouterFns, func(rGroup *gin.RouterGroup) {
		roleRouter(rGroup,handler.NewRoleHandler())
	})
}

func roleRouter(group *gin.RouterGroup, h handler.RoleHandler)  {
	g := group.Group("/role")
	g.POST("/create", h.Create)
	g.PUT("/update/:id", h.UpdateByID)
	g.DELETE("/delete/:id", h.DeleteByID)
	g.GET("/list", h.GetColumn)
}
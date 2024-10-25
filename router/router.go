package router

import (
	"HiChat/service"
	"github.com/gin-gonic/gin"
)

func Roter() *gin.Engine {
	//初始化路由
	router := gin.Default()

	//v1版本
	v1 := router.Group("v1")

	//用户模块，后续有个用户的api就放置其中
	user := v1.Group("user")
	{
		user.GET("/list", service.List)
	}

	return router

}

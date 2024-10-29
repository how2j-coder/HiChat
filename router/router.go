package router

import (
	result "HiChat/common"
	"HiChat/global"
	"HiChat/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, result.NotFound)
}

func Roter() *gin.Engine {
	//初始化路由
	router := gin.Default()

	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		global.Logger.Error(err.Error())
	}

	router.NoRoute(NotFound)

	//v1版本
	v1 := router.Group("v1")

	//用户模块，后续有个用户的api就放置其中
	user := v1.Group("user")
	{
		user.POST("/create", service.Create) // 创建用花
		user.GET("/list", service.List)
	}

	return router
}

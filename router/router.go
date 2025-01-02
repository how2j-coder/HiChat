package router

import (
	result "HiChat/common"
	"HiChat/global"
	"HiChat/middlewear"
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

	// 注册自定义的参数效验
	//utils.GisterValidation()


	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		global.Logger.Error(err.Error())
	}

	router.NoRoute(NotFound)

	api := router.Group("/api")

	//v1版本
	v1 := api.Group("v1")

	//用户模块，后续有个用户的api就放置其中
	user := v1.Group("user")
	{
		user.POST("/create", service.Create) // 创建用户
		user.POST("/login", service.LoginByPassword)
		user.PATCH("/update", middlewear.AuthRequired(), service.UpDateUser)
		user.GET("/list", middlewear.AuthRequired(), service.List)
		user.DELETE("/del", middlewear.AuthRequired(), service.DeleteUser)
		user.PATCH("/restore", service.RestoreUser)
	}

	//文件模块
	file := v1.Group("/file")
	{
		file.POST("/upload", middlewear.AuthRequired(), service.FileUploadSingle)
		file.POST("/uploads", middlewear.AuthRequired(), service.FileUploadMultiple)
		file.GET("/download/:fileName", service.FileDownload)
	}

	//系统平台
	platform := v1.Group("platform")
	{
		platform.POST("/create", service.CratePlatform)
		platform.PATCH("/update",  service.UpdatePlatform)
		platform.GET("/list", service.FindPlatformList)
		platform.DELETE("/del/:id", service.DeletePlatform)
	}

	//菜单
	menus := v1.Group("menus")
	{
		menus.GET("/tree", service.GetMenuTree)
		menus.POST("/create", service.CreateMenu)
		menus.PATCH("/update", service.UpdateMenu)
	}

	return router
}

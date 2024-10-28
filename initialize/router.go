package initialize

import (
	"HiChat/global"
	"HiChat/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	route := router.Roter()
	err := route.Run(fmt.Sprintf(":%d", global.ServiceConfig.Port))
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return route
}

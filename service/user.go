package service

import (
	result "HiChat/common"
	"HiChat/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(ctx *gin.Context) {
	list, err := dao.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusOK, result.Success.WithData(list))
	}
}

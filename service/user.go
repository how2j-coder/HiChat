package service

import (
	result "HiChat/common"
	"HiChat/dao"
	"HiChat/global"
	"HiChat/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

// Create 创建用户
func Create(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Request.FormValue("username") // 用户名称
	//pasword := ctx.Request.FormValue("password")
	////生成盐值
	//salt := fmt.Sprintf("%d", rand.Int31())
	rand.Int31()
	//user.PassWord = common.SaltPassWord(password, salt)
	//user.Salt = salt
	cUser, err := dao.CreateUser(user)
	if err != nil {
		return
	} else {
		ctx.JSON(200, result.Success.WithData(cUser))
	}

}

// List 获取用户列表
func List(ctx *gin.Context) {
	list, err := dao.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Error.WithMsg(err.Error()))
		global.Logger.Error("Get user list error", zap.Error(err))
	} else {
		if list == nil {
			list = make([]*models.UserBasic, 0)
		}
		ctx.JSON(http.StatusOK, result.Success.WithData(list))
	}

}

package service

import (
	. "HiChat/common"
	"HiChat/dao"
	"HiChat/global"
	"HiChat/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

// Create 创建用户
func Create(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Request.FormValue("username") // 用户名称
	password := ctx.Request.FormValue("password")
	//生成盐值
	salt := fmt.Sprintf("%d", rand.Int31())
	user.PassWord = SaltPassWord(password, salt)
	user.Salt = salt
	test, _ := json.MarshalIndent(user, "", " ")
	global.Logger.Info(string(test))
	//cUser, err := dao.CreateUser(user)
	ctx.JSON(200, Success.WithData("12313"))
	//if err != nil {
	//	return
	//} else {
	//	ctx.JSON(200, Success.WithData(cUser))
	//}

}

// List 获取用户列表
func List(ctx *gin.Context) {
	list, err := dao.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
		global.Logger.Error("Get user list error", zap.Error(err))
	} else {
		if list == nil {
			list = make([]*models.UserBasic, 0)
		}
		ctx.JSON(http.StatusOK, Success.WithData(list))
	}

}

package service

import (
	. "HiChat/common"
	"HiChat/dao"
	"HiChat/global"
	"HiChat/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.?/|*&$%#@!{}[]"

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Create 创建用户
func Create(ctx *gin.Context) {
	type TempData struct {
		Name     string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	temp := TempData{}
	user := models.UserBasic{}
	if err := ctx.ShouldBind(&temp); err != nil {
		ctx.JSON(http.StatusBadRequest, ParamsError.WithMsg(err.Error()))
		global.Logger.Error(err.Error())
		return
	}
	user.Name = temp.Name
	user.Email = temp.Email
	password := temp.Password // 密码

	if temp.Name == "" {
		user.Name = generateRandomString(7)
	}

	if user.Email == "" || password == "" {
		ctx.JSON(http.StatusOK, ParamsError.WithMsg("邮箱或密码不能为空！"))
		return
	}

	findUser, err := dao.FindUser(user)

	if (err == nil && findUser != nil) || err != nil {
		ctx.JSON(http.StatusOK, ParamsError.WithMsg("用户已注册！"))
		return
	}

	//生成盐值
	salt := fmt.Sprintf("%d", rand.Int31())
	user.PassWord = SaltPassWord(password, salt)
	user.Salt = salt
	
	_, err = dao.CreateUser(user)
	if err != nil {
		return
	} else {
		ctx.JSON(200, Success.WithMsg("注册成功！"))
	}
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

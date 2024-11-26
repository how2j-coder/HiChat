package service

import (
	. "HiChat/common"
	"HiChat/dao"
	"HiChat/global"
	"HiChat/models"
	"HiChat/utils"
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
		Email    string `json:"email" binding:"required" requiredMsg:"邮箱不能为空"`
		Password string `json:"password" binding:"required" requiredMsg:"密码不能为空"`
	}
	temp := TempData{}
	user := models.User{}
	if err := ctx.ShouldBind(&temp); err != nil {
		errText := utils.GetErrorMsg(err, temp)
		ctx.JSON(http.StatusBadRequest, ParamsNilError.WithMsg(errText))
		global.Logger.Error(err.Error())
		return
	}
	user.Name = temp.Name
	user.Email = temp.Email
	password := temp.Password // 密码

	if temp.Name == "" {
		user.Name = generateRandomString(7)
	}

	findUser, err := dao.FindUser(user)

	if err == nil && findUser != nil {
		ctx.JSON(http.StatusOK, ParamsNilError.WithMsg("用户已注册！"))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ParamsNilError.WithMsg(err.Error()))
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

// Login 用户登录
func Login(ctx *gin.Context) {
	type TempData struct {
		Name     string `json:"username" binding:"required" requiredMsg:"用户名不能为空"`
		Password string `json:"password"`
	}
	temp := TempData{}
	if err := ctx.ShouldBind(&temp); err != nil {
		errText := utils.GetErrorMsg(err, temp)
		ctx.JSON(http.StatusBadRequest, ParamsNilError.WithMsg(errText))
		global.Logger.Error(err.Error())
		return
	}

	ctx.JSON(200, Success.WithData("test"))
}

// List 获取用户列表
func List(ctx *gin.Context) {
	list, err := dao.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
		global.Logger.Error("Get user list error", zap.Error(err))
	} else {
		if list == nil {
			list = make([]*models.User, 0)
		}
		ctx.JSON(http.StatusOK, Success.WithData(list))
	}

}

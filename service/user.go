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
	"time"
)



// Create 创建用户
func Create(ctx *gin.Context) {
	type TempData struct {
		Name     string `json:"username"`
		Email    string `json:"email" binding:"required" requiredMsg:"邮箱不能为空"`
		Password string `json:"password" binding:"required,min=6" requiredMsg:"密码不能为空" minMsg:"密码最少为6位数"`
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
		user.Name = utils.GenerateRandomString(7)
	}

	findUser, err := dao.FindUser(user)

	if err == nil && findUser != nil {
		ctx.JSON(http.StatusOK, Success.WithMsg("用户已注册"))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
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
		ctx.JSON(http.StatusOK, Success.WithMsg("注册成功"))
	}
}

// LoginByPassword 用户登录
func LoginByPassword(ctx *gin.Context) {
	type TempData struct {
		Name     string `json:"username" binding:"required" requiredMsg:"用户名不能为空"`
		Password string `json:"password" binding:"required" requiredMsg:"密码不能为空"`
	}
	temp := TempData{}
	if err := ctx.ShouldBind(&temp); err != nil {
		errText := utils.GetErrorMsg(err, temp)
		ctx.JSON(http.StatusBadRequest, ParamsNilError.WithMsg(errText))
		global.Logger.Error(err.Error())
		return
	}

	findUser, err := dao.FindUserByName(temp.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg("登录失败"))
		global.Logger.Error(err.Error())
		return
	}

	if findUser == nil {
		ctx.JSON(http.StatusOK, NotFound.WithMsg("用户未注册"))
		return
	}

	checkPwdOk := CheckPassWord(temp.Password, findUser.Salt, findUser.PassWord)

	if !checkPwdOk {
		ctx.JSON(http.StatusOK, Error.WithMsg("密码错误"))
		return
	}

	acToken, reToken := GenerateTaken(findUser.ID,"Auth_Server", time.Hour * 2)
	ctx.JSON(200, Success.WithData(map[string]string{
		"access_token": acToken,
		"refresh_token":reToken,
	}))
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

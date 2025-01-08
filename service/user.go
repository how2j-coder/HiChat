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

	findUser, err := dao.FindUserByEmail(user.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
		return
	}

	if  findUser != nil {
		ctx.JSON(http.StatusOK, Success.WithMsg("用户已注册"))
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
		Account     string `json:"account" binding:"required" requiredMsg:"账号不能为空"`
		Password string `json:"password" binding:"required" requiredMsg:"密码不能为空"`
	}
	temp := TempData{}
	if err := ctx.ShouldBind(&temp); err != nil {
		errText := utils.GetErrorMsg(err, temp)
		ctx.JSON(http.StatusOK, ParamsNilError.WithMsg(errText))
		return
	}

	findUser, err := dao.FindUserByEmail(temp.Account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg("登录失败"))
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

// DeleteUser 删除用户
func DeleteUser(ctx *gin.Context)  {
	userId, _ := ctx.Get(global.AuthCtxFiled)
	findUser, err := dao.FindUserById(userId.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
		return
	}
	if findUser == nil {
		ctx.JSON(http.StatusOK, NotFound.WithMsg("未查询到用户"))
		return
	}
	err = dao.DeleteUser(*findUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, Success)
}

// RestoreUser 用户数据恢复
func RestoreUser(ctx *gin.Context)  {
	email := ctx.Query("email")
	findUser, err := dao.FindUserById(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
		return
	}
	if findUser == nil {
		ctx.JSON(http.StatusOK, NotFound.WithMsg("未查询到用户"))
		return
	}
	user, err := dao.UnDeleteUser(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, Success.WithData(user))
}

// UpDateUser 更新用户数据
func UpDateUser(ctx *gin.Context)  {
	type TempUser struct {
		Avatar string `json:"avatar"`
		Gender string `json:"gender"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	}
	userId, _ := ctx.Get(global.AuthCtxFiled)
	findUser, err := dao.FindUserById(userId.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
		return
	}
	if findUser == nil {
		ctx.JSON(http.StatusOK, NotFound.WithMsg("未查询到用户"))
		return
	}

	tempUser := TempUser{}
	if err := ctx.ShouldBind(&tempUser); err != nil {
		errText := utils.GetErrorMsg(err, tempUser)
		ctx.JSON(http.StatusBadRequest, ParamsNilError.WithMsg(errText))
		return
	}

	findUser.Avatar = tempUser.Avatar
	findUser.Gender = tempUser.Gender
	findUser.Phone = tempUser.Phone
	findUser.Email = tempUser.Email

	_, err = dao.UpdateUser(*findUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error.WithMsg(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, Success.WithData(tempUser))
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

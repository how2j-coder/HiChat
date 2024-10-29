package test

import (
	"HiChat/dao"
	"HiChat/initialize"
	"HiChat/models"
	"testing"
)

func setup() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
}

// 创建用户
func TestCreateUser(t *testing.T) {
	setup()
	user := models.UserBasic{
		Name:     "how2j",
		PassWord: "how2j.online",
		Avatar:   "https://example.com",
		Gender:   "male",
		Phone:    "15683832914",
		Email:    "how2j@gmail.com",
	}
	result, err := dao.CreateUser(user)

	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(result.ID, user.ID)
	}
}

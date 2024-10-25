package test

import (
	"HiChat/dao"
	"HiChat/initialize"
	"HiChat/models"
	"encoding/json"
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

// 删除用户
func TestDeleteUser(t *testing.T) {
	setup()
	user, err := dao.GetUser(models.UserBasic{
		Phone: "15683832914",
		Email: "how2j@gmail.com",
	})

	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("find user success", user)
	}

	err = dao.DeleteUser(*user)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("delete user", user)
	}
}

// 恢复用户
func TestUnDeleteUser(t *testing.T) {
	setup()
	findUser, err := dao.GetUser(models.UserBasic{
		Phone: "15683832914",
		Email: "how2j@gmail.com",
	})
	if err != nil {
		t.Error(err.Error())
	}
	if findUser == nil {
		user, err := dao.UnDeleteUser(models.UserBasic{
			Phone: "15683832914",
			Email: "how2j@gmail.com",
		})
		if err != nil {
			t.Error(err.Error())
		} else {
			t.Log("unDelete user success", user)
		}
	}
}

// 获取用户列表
func TestGetUserList(t *testing.T) {
	setup()
	userList, err := dao.GetUserList()
	if err != nil {
		t.Error(err.Error())
	} else {
		user, _ := json.MarshalIndent(userList, "", " ")
		str := string(user)
		t.Log("find user list", str)
	}
}

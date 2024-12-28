package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// CreateUser 创建用户
func CreateUser(user models.User) (*models.User, error) {
	tx := global.DB.Create(&user)
	if tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, errors.New("新建用户失败")
	}
	return &user, nil
}

// UpdateUser 更新数据
func UpdateUser(user models.User) (*models.User, error) {
	fmt.Println(user)
	tx := global.DB.Model(&user).Select("Avatar",
		"Gender", "Phone", "Email",
	).Updates(user)

	if tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &user, nil
}

// FindUserByName 通过用户名精准查询
func FindUserByName(name string) (*models.User, error) {
	user := models.User{}
	if tx := global.DB.Where("name = ?", name).First(&user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &user, nil
}

// FindUserByEmail  查找用户-精准查询(根据email)
func FindUserByEmail(user models.User) (*models.User, error) {
	if tx := global.DB.Where("email = ?", user.Email).First(&user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &user, nil
}

// FindUserById  查找用户-精准查询(根据userId)
func FindUserById(userID string) (*models.User, error)  {
	user := models.User{}
	if tx := global.DB.Where("id = ?", userID).First(&user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &user, nil
}

// DeleteUser 删除用户
func DeleteUser(user models.User) error {
	if tx := global.DB.Delete(&user); tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return errors.New("删除用户失败")
	}
	return nil
}

// UnDeleteUser 恢复用户 (根据 email)
func UnDeleteUser(email string) (*models.User, error) {
	user := models.User{}
	if tx := global.DB.Unscoped().Model(&user).Where("email = ?", email).First(&user).Update("deleted_at", nil); tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, errors.New("用户找回失败")
	}

	return &user, nil
}

// GetUserList 获取用户列表
func GetUserList() ([]*models.User, error) {
	var list []*models.User
	if tx := global.DB.Omit("Salt", "Identity",
		"PassWord", "ClientIp", "ClientPort", "UpdatedAt", "DeletedAt").Find(&list); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return list, nil
}

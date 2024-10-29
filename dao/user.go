package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"gorm.io/gorm"
)

// CreateUser 创建用户
func CreateUser(user models.UserBasic) (*models.UserBasic, error) {
	tx := global.DB.Create(&user)
	if tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, errors.New("新建用户失败")
	}
	return &user, nil
}

// FindUserByName 通过用户名精准查询
func FindUserByName(name string) (*models.UserBasic, error) {
	user := models.UserBasic{}
	if tx := global.DB.Where("name = ?", name).First(&user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &user, nil
}

// FindUser 查找用户-精准查询(根据phone Or email)
func FindUser(user models.UserBasic) (*models.UserBasic, error) {
	if tx := global.DB.Where("phone = ?", user.Phone).Or("email = ?", user.Email).First(&user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &user, nil
}

// DeleteUser 删除用户
func DeleteUser(user models.UserBasic) error {
	if tx := global.DB.Delete(&user); tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return errors.New("删除用户失败")
	}
	return nil
}

// UnDeleteUser 恢复用户 (根据phone Or email)
func UnDeleteUser(user models.UserBasic) (*models.UserBasic, error) {
	if tx := global.DB.Unscoped().Model(&user).Where(
		"phone = ?", user.Phone,
	).Or("email = ?", user.Email).First(&user).Update("deleted_at", nil); tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, errors.New("用户找回失败")
	}

	return &user, nil
}

// GetUserList 获取用户列表
func GetUserList() ([]*models.UserBasic, error) {
	var list []*models.UserBasic
	if tx := global.DB.Find(&list); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return list, nil
}

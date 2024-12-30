package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"gorm.io/gorm"
)

// CratePlatform 创建系统平台
func CratePlatform(platform models.Platform) (*models.Platform, error)  {
	tx := global.DB.Create(&platform)
	if tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &platform, nil
}

// FindNameToPlatform 根据名称查找
func FindNameToPlatform(platformName string) (*models.Platform, error)  {
	platform := models.Platform{}
	tx := global.DB.Where("platform_name = ?", platformName).First(&platform)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &platform, nil
}

// FindIdToPlatform 根据Id查找
func FindIdToPlatform(platformId string) (*models.Platform, error)  {
	platform := models.Platform{}
	tx := global.DB.Where("id = ?", platformId).First(&platform)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &platform, nil
}

// FindPlatformList 查询平台信息
func FindPlatformList() ([]*models.Platform, error)  {
	var platforms []*models.Platform
	if tx := global.DB.Find(&platforms); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return platforms, nil
}

// UpdatePlatform 更新系统平台信息
func UpdatePlatform(id string, data map[string]interface{}) (*models.Platform, error)  {
	platform := models.Platform{}
	platform.ID = id
	tx := global.DB.Model(&platform).Omit("PlatformCode").Updates(data).First(&platform)
	if tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &platform, nil
}
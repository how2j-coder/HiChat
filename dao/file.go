package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"gorm.io/gorm"
)

// UploadSingle 上传单个文件
func UploadSingle(file models.File) (*models.File, error) {
	tx := global.DB.Create(&file)
	if tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &file, nil
}

// FindFileFileName 查找文件
func FindFileFileName(fileName string) (*models.File, error)  {
	file := models.File{}
	tx := global.DB.Where("file_name = ?", fileName).First(&file)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &file, nil
}
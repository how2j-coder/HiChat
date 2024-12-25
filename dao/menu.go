package dao

import (
	"HiChat/global"
	"HiChat/models"
)

// CreateMenu 菜单创建
func CreateMenu(menu models.Menu) (*models.Menu, error)  {
	tx := global.DB.Create(&menu)
	if tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &menu, nil
}
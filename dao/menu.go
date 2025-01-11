package dao

import (
	"HiChat/global"
	"HiChat/models"
	"errors"
	"gorm.io/gorm"
)

// CreateMenu 菜单创建
func CreateMenu(menu models.Menu) (*models.Menu, error)  {
	var menus []models.Menu
	countTx := global.DB.Where(
		"platform_id = ? AND parent_menu_id = ?",
		menu.PlatformID, menu.ParentMenuID,
		).Find(&menus)
	if countTx.Error != nil {
		global.Logger.Error(countTx.Error.Error())
		return nil, countTx.Error
	}

	sort := int(countTx.RowsAffected)
	menu.SortOrder = &sort
	tx := global.DB.Create(&menu)
	if tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &menu, nil
}

// FindMenuCodeToMenu 根据 Menu Code 查找菜单
func FindMenuCodeToMenu(menuCode string) (*models.Menu, error)  {
	menu :=  models.Menu{}
	tx := global.DB.Where("menu_code = ?", menuCode).First(&menu)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &menu, nil
}

// FindIdToMenu 根据ID查询菜单信息
func FindIdToMenu(menuId string) (*models.Menu, error)  {
	menu :=  models.Menu{}
	tx := global.DB.Where("id = ?", menuId).First(&menu)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &menu, nil
}

// UpdateMenuIdToMenu 更新菜单数据
func UpdateMenuIdToMenu(menuId string, data map[string]interface{}) (*models.Menu, error)  {
	menu :=  models.Menu{}
	menu.ID = menuId
	tx := global.DB.Model(&menu).Select(
		"PlatformID", "ParentMenuID", "MenuName",
	"MenuCode", "MenuType", "MenuFilePath", "IsVisible",
	"IsEnabled", "IsRefresh", "SortOrder",
		).Updates(data)
	if tx.Error != nil {
		global.Logger.Error(tx.Error.Error())
		return nil, tx.Error
	}
	return &menu, nil
}

// FindPlatformToMenus 获取平台对应的菜单数据
func FindPlatformToMenus(platformID string, menuID string) ([]*models.Menu, error)  {
	var menu []*models.Menu
	if menuID == "" {
		tx := global.DB.Where("platform_id = ?", platformID).Find(&menu)
		if tx.Error != nil {
			if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, tx.Error
		}
	} else {
		tx := global.DB.Where(
			"platform_id = ? AND parent_menu_id = ?",
			platformID,
			menuID,
			).Find(&menu)
		if tx.Error != nil {
			if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, tx.Error
		}
	}


	return menu, nil
}

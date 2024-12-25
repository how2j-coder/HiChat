package models

type Menu struct {
	CommonBaseModel
	PlatformID string `json:"platform_id"`
	Platform Platform `gorm:"foreignkey:PlatformID" json:"-"`
	ParentMenuID int `gorm:"comment:父级菜单ID"`
	MenuName string `gorm:"comment:菜单名称"`
	MenuCode string `gorm:"comment:菜单Code"`
	MenuType int `gorm:"comment:菜单类型(0 目录 1 菜单 2 按钮)"`
	MenuFilePath string `gorm:"comment:菜单对应的前端文件模板路径"`
	IsVisible int `gorm:"comment:是否可见(0 隐藏 1 显示)"`
	IsEnabled int `gorm:"comment:是否启用(0 关闭 1 启用)"`
	IsRefresh int `gorm:"comment:打开页面时是否刷新页面(0 否 1 是)"`
	SortOrder int  `gorm:"comment:菜单顺序"`
}

func (table *Menu) MenuTableName() string  {
	return "menus"
}
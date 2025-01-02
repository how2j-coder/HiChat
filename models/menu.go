package models

type Menu struct {
	CommonBaseModel
	PlatformID string `json:"platform_id" binding:"required" requiredMsg:"请选择平台"`
	Platform Platform `gorm:"foreignkey:PlatformID" json:"-"`
	ParentMenuID string `gorm:"comment:父级菜单ID" json:"parent_menu_id"`
	MenuName string `gorm:"comment:菜单名称" json:"menu_name" binding:"required" requiredMsg:"菜单名称不能为空"`
	MenuCode string `gorm:"comment:菜单Code" json:"menu_code" binding:"required" requiredMsg:"菜单Code不能为空"`
	MenuType *int `gorm:"default:1;comment:菜单类型(0 目录 1 菜单 2 按钮)" json:"menu_type"`
	MenuPath string `gorm:"comment:菜单路由地址" json:"menu_path" binding:"required"`
	MenuFilePath string `gorm:"comment:菜单对应的前端文件模板路径" json:"menu_file_path" binding:"required" requiredMsg:"模板路径不能为空"`
	IsVisible *int `gorm:"default:1;comment:是否可见(0 隐藏 1 显示)" json:"is_visible"`
	IsEnabled *int `gorm:"default:1;comment:是否启用(0 关闭 1 启用)" json:"is_enabled"`
	IsRefresh *int `gorm:"default:0;comment:打开页面时是否刷新页面(0 否 1 是)" json:"is_refresh"`
	SortOrder *int  `gorm:"comment:菜单顺序" json:"sort_order"`
	Children []*Menu `gorm:"-" json:"children"`
}

//func (table *Menu) UnmarshalJSON(data []byte) error {
//	type Alias Menu
//	alias := &struct {
//		ParentMenuID string `json:"parent_menu_id"`
//		*Alias
//	}{
//		Alias: (*Alias)(table),
//	}
//	if err := json.Unmarshal(data, alias); err != nil {
//		return err
//	}
//	if string(alias.ParentMenuID) != ""  && len(alias.ParentMenuID) > 0 {
//		table.ParentMenuID = sql.NullString{String: alias.ParentMenuID, Valid: true}
//	} else {
//		table.ParentMenuID = sql.NullString{Valid: false}
//	}
//	return nil
//}

func (table *Menu) MenuTableName() string  {
	return "menus"
}
package model

import "com/chat/service/pkg/datastore/mysql"

type Menu struct {
	mysql.BaseModel `gorm:"embedded"`
	PlatformID uint64 `gorm:"comment:关联的平台"`
	ParentMenuID    uint64 `gorm:"default:0;comment:父级菜单ID"`
	MenuName        string `gorm:"type:varchar(20);comment:菜单名称"`
	MenuCode        string `gorm:"type:varchar(20);unique_index;not null;comment:菜单唯一Code"`
	MenuPath        string `gorm:"type:varchar(255);not null;comment:菜单对于地址"`
	MenuSource      string `gorm:"type:varchar(255);not null;comment:菜单对应前端的文件页面地址"`
	MenuIcon        string `gorm:"type:varchar(100);comment:菜单图标"`
	Sort            int    `gorm:"type:int;comment:菜单排序"`
	IsEnable        int8   `gorm:"default:1;not null;comment:菜单状态（1：启用，0：禁用）"`
	Type            int8   `gorm:"default:1;not null;comment:类型（1：菜单，2：按钮，3：目录）"`
	IsRefresh       int8   `gorm:"default:1;not null;comment:页面刷新（1：刷新，0：不刷新）"`
	IsVisible       int8   `gorm:"default:1;comment:是否可见(0 隐藏 1 显示)"`
	IsSingle        int8   `gorm:"default:0;comment:存在一个子级菜单时是否合并显示路由"`
	Role            []Role `gorm:"many2many:role_menu" json:"-"`
}

func (u *Menu) TableName() string {
	return "menu"
}

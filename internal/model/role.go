package model

import (
	"com/chat/service/pkg/datastore/mysql"
	"com/chat/service/pkg/utils"
	"gorm.io/gorm"
)

type Role struct {
	mysql.BaseModel `gorm:"embedded"`
	ParentRoleID uint64 `gorm:"column:parent_role_id"`
	RoleName string `gorm:"type:varchar(56);unique_index;not null"`
	Remark string `gorm:"type:varchar(255);"`
	User []User `gorm:"many2many:role_user;"`
	Menu []Menu `gorm:"many2many:role_menu;"`
}

func (r *Role) TableName() string {
	return "roles"
}


var s, _ = utils.NewSnowflake(1)


// RoleUser user 和 role 的连接表
type RoleUser struct {
	ID        uint64         `gorm:"column:id;primary_key" json:"id"`
	UserID uint64 `gorm:"column:user_id;foreignKey:UserID;references:users(id)"`
	RoleID uint64 `gorm:"column:role_id;foreignKey:RoleID;references:roles(id)"`
}

func (r *RoleUser) TableName() string {
	return "role_user"
}

func (r *RoleUser) BeforeCreate(_ *gorm.DB) (err error)  {
	generate, err := s.Generate()
	r.ID = uint64(generate)
	return
}

// RoleMenu user 和 menu 的连接表
type RoleMenu struct {
	ID        uint64         `gorm:"column:id;primary_key" json:"id"`
	MenuID uint64 `gorm:"column:menu_id;foreignKey:MenuID;references:menu(id)"`
	RoleID uint64 `gorm:"column:role_id;foreignKey:RoleID;references:roles(id)"`
}

func (r *RoleMenu) TableName() string {
	return "role_menu"
}

func (r *RoleMenu) BeforeCreate(_ *gorm.DB) (err error)  {
	generate, err := s.Generate()
	r.ID = uint64(generate)
	return
}

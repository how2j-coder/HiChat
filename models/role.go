package models

type Role struct {
	CommonBaseModel
	RoleName string `gorm:"comment:角色名称"`
	IsEnabled *int `gorm:"comment:是否启用"`
}

func (Role) TableName() string {
	return "role"
}

type RoleMenu struct {
	CommonBaseModel
	RoleId string `gorm:"comment:角色Id"`
	Role Role   `gorm:"foreignKey:RoleId" json:"-"`
	MenuId string `gorm:"comment:菜单Id"`
	Menu   Menu   `gorm:"foreignKey:MenuId" json:"-"`
}

func (RoleMenu) TableName() string {
	return "role_menu"
}

type RoleUser struct {
	CommonBaseModel
	RoleId string `json:"-"`
	Role   Role   `gorm:"foreignKey:RoleId" json:"-"`
	UserId string `json:"-"`
	User   User   `gorm:"foreignKey:UserId" json:"-"`
}

func (RoleUser) TableName() string {
	return "role_user"
}
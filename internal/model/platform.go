package model

import "com/chat/service/pkg/datastore/mysql"

type Platform struct {
	mysql.BaseModel `gorm:"embedded"`
	PlatformName  string `gorm:"type:varchar(255);unique_index;not null"`
	PlatformCode  string `gorm:"type:varchar(255);unique_index;not null"`
	IsEnabled     int8   `gorm:"default:1;comment:是否启用(0 停用 1 启用)"`
	PlatformUrl   string `gorm:"type:varchar(255);not null"`
	Menus []Menu `gorm:"foreignkey:PlatformID;"`
}

func (p *Platform) TableName() string {
	return "platform"
}
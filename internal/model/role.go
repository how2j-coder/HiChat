package model

import "com/chat/service/pkg/datastore/mysql"

type Role struct {
	mysql.BaseModel `gorm:"embedded"`
	RoleName string `gorm:"type:varchar(56);unique_index;not null"`
	Remark string `gorm:"type:varchar(255)"`
}

func (r *Role) TableName() string {
	return "role"
}
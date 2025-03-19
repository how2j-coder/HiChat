package model

import "com/chat/service/pkg/datastore/mysql"

type User struct {
	mysql.BaseModel `gorm:"embedded"`
	Username        string `gorm:"type:varchar(255);"`
	Account         string `gorm:"type:varchar(6);uniqueIndex:idx_username_password;not null"`
	Password        string `gorm:"type:varchar(255)"`
	Email           string `gorm:"type:varchar(100);uniqueIndex:idx_username_password;unique"`
	AvatarURL       string `gorm:"type:varchar(255)"`
	Gender          string `gorm:"type:varchar(10);check:gender IN ('Male', 'Female', 'Other');default:'Other'"`
	Type            string `gorm:"type:varchar(10);check:type IN ('Admin', 'Ordinary');default:'ordinary'"`
	Role            []Role `gorm:"many2many:role_user"`
}

func (u *User) TableName() string {
	return "users"
}

package model

import "com/chat/service/pkg/datastore/mysql"

type User struct{
	mysql.BaseModel `gorm:"embedded"`
	Name string `gorm:"type:varchar(255);uniqueIndex:idx_username_password;not null"`
	PasswordHash string `gorm:"type:varchar(255)"`
	Email       string `gorm:"type:varchar(100);uniqueIndex:idx_username_password;unique"`
	AvatarURL   string `gorm:"type:varchar(255)"`
	Gender      string `gorm:"type:varchar(10);check:gender IN ('Male', 'Female', 'Other');default:'Other'"`
}

func (u *User) TableName() string {
	return "users"
}

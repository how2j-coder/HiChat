package model

import "com/chat/service/pkg/datastore/mysql"

type Upload struct {
	mysql.BaseModel  `gorm:"embedded"`
	Type             string `gorm:"type:varchar(32);not null"`
	MimeType         string `gorm:"type:varchar(128);not null"`
	Size             int64  `gorm:"not null;comment:文件大小(字节)"`
	MD5              string `gorm:"type:varchar(32);not null;index;comment:文件MD5值"`
	Key              string `gorm:"type:varchar(64);not null;uniqueIndex;comment:存储服务中的唯一标识(如OSS的object key)"`
	Storage          string `gorm:"type:varchar(20);not null;comment:存储方式(local/oss/cos/s3)"`
	OriginalFileName string `gorm:"type:varchar(255);not null;comment:文件名称"`
	RelativeFile     string `gorm:"type:varchar(255);not null;comment:文件存储地址"`
	URL              string `gorm:"type:varchar(512);not null;comment:访问URL"`
	Status           int8   `gorm:"type:tinyint;default:1;comment:状态(0=禁用,1=正常)"`
}

func (f *Upload) TableName() string {
	return "upload"
}
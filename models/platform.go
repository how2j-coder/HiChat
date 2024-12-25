package models

type Platform struct {
	CommonBaseModel
	PlatformName string `gorm:"comment:平台名称"`
	PlatformCode string `gorm:"comment:平台Code"`
	PlatformUrl string `gorm:"comment:平台地址"`
	Version string `gorm:"comment:平台版本"`
	IsEnable int `gorm:"comment:是否启用"`
}

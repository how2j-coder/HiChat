package models

type Platform struct {
	CommonBaseModel
	PlatformName string `gorm:"comment:平台名称" json:"name"`
	PlatformCode string `gorm:"comment:平台Code" json:"code"`
	PlatformUrl string `gorm:"comment:平台地址" json:"url"`
	Version string `gorm:"comment:平台版本" json:"version"`
	IsEnable int `gorm:"comment:是否启用(1 启用 0 停用);default:1" json:"is_enable"`
}

func (table *Platform) TableName() string {
	return "platform"
}

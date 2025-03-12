package model

type CasbinRule struct {
	Ptype string `gorm:"type:varchar(10)" json:"ptype"`
	V0    string `gorm:"type:varchar(255)" json:"v0"`
	V1    string `gorm:"type:varchar(255)" json:"v1"`
	V2    string `gorm:"type:varchar(255)" json:"v2"`
	V3    string `gorm:"type:varchar(255)" json:"v3"`
	V4    string `gorm:"type:varchar(255)" json:"v4"`
	V5    string `gorm:"type:varchar(255)" json:"v5"`
}

func (r *CasbinRule) TableName() string {
	return "casbin_rule"
}


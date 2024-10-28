package models

import (
	"HiChat/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type CommonBaseModel struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var snowflake = utils.NewSnowflake(int64(52))

func (com *CommonBaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	com.ID = strconv.FormatInt(snowflake.GenerateID(), 10)
	return nil
}

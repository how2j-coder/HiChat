package models

import (
	"HiChat/utils"
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t *LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(*t)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *LocalTime) IsZero() bool {
	tTime := time.Time(*t)
	return tTime.IsZero()
}

func (t *LocalTime) Format(format string) string {
	tTime := time.Time(*t)
	return tTime.Format(format)
}

type CommonBaseModel struct {
	ID        string `gorm:"primary_key" json:"userId"`
	CreatedAt *LocalTime			`json:"created_at,omitempty"`
	UpdatedAt *LocalTime      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

var snowflake = utils.NewSnowflake(int64(52))

func (com *CommonBaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	com.ID = strconv.FormatInt(snowflake.GenerateID(), 10)
	return nil
}

// AfterFind 钩子函数：在查询后处理返回的结果
//func (com *CommonBaseModel) AfterFind(tx *gorm.DB) (err error) {
//	// 如果 UpdatedAt 是零值，则将其设置为 null
//	if com.UpdatedAt.IsZero() {
//	}
//	// 如果 DeletedAt 是零值，则将其设置为 nil
//	//if com.DeletedAt.Time.IsZero() {
//	//	com.DeletedAt = &gorm.DeletedAt{} // 设置为零值
//	//}
//	return nil
//}

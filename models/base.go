package models

import (
	"HiChat/utils"
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type localTime time.Time

func (t *localTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t *localTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(*t)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *localTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = localTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *localTime) IsZero() bool {
	tTime := time.Time(*t)
	return tTime.IsZero()
}

func (t *localTime) Format(format string) string {
	tTime := time.Time(*t)
	return tTime.Format(format)
}

type commonBaseModel struct {
	ID        string `gorm:"primary_key" json:"id"`
	CreatedAt *localTime			`json:"created_at,omitempty"`
	UpdatedAt *localTime      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

var snowflake = utils.NewSnowflake(int64(52))

func (com *commonBaseModel) BeforeCreate(_ *gorm.DB) (err error) {
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

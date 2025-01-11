package utils

import (
	"gorm.io/gorm"
)
type PageDto struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}


func Paginate(req *PageDto) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := req.Page
		if page == 0 {
			page = 1
		}
		pageSize := req.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
			case pageSize <= 0:
				pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
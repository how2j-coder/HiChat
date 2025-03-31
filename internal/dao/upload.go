package dao

import (
	"com/chat/service/internal/model"
	"context"
	"gorm.io/gorm"
)

type UploadDao interface {
	CreateUpload(c context.Context, upload *model.Upload) error
}

type uploadDao struct {
	db *gorm.DB
}

var _ UploadDao = (*uploadDao)(nil)

func NewUploadDao(db *gorm.DB) UploadDao {
	return &uploadDao{db: db}
}

func (u *uploadDao) CreateUpload(c context.Context, upload *model.Upload) error {
	return u.db.WithContext(c).Create(upload).Error
}

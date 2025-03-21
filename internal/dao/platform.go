package dao

import (
	"com/chat/service/internal/model"
	"context"
	"gorm.io/gorm"
)

type PlatformDao interface {
	Create(ctx context.Context, table *model.Platform) error
	UpdateByID(ctx context.Context, table *model.Platform, update map[string]interface{}) error
	DeleteByID(ctx context.Context, table *model.Platform) error
	GetColumn(ctx context.Context) ([]*model.Platform, error)
}

type platformDao struct {
	db *gorm.DB
}

var _ PlatformDao = (*platformDao)(nil)

func NewPlatformDao(db *gorm.DB) PlatformDao {
	return &platformDao{
		db: db,
	}
}

func (d *platformDao) Create(ctx context.Context, table *model.Platform) error {
	return d.db.WithContext(ctx).Create(&table).Error
}

func (d *platformDao) UpdateByID(ctx context.Context, table *model.Platform, update map[string]interface{}) error {
	err := d.db.WithContext(ctx).Model(table).Updates(update).Error
	return err
}

func (d *platformDao) DeleteByID(ctx context.Context, table *model.Platform) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(table).Error; err != nil {
			return err
		}
		return nil
	})
}

func (d *platformDao) GetColumn(ctx context.Context) ([]*model.Platform, error) {
	var platforms []*model.Platform
	if err := d.db.WithContext(ctx).Find(&platforms).Error; err != nil {
		return nil, err
	}
	return platforms, nil
}
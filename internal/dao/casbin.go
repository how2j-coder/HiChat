package dao

import (
	"com/chat/service/internal/model"
	"context"
	"gorm.io/gorm"
)

type CasbinRuleDao interface {
	Create(ctx context.Context, table *model.CasbinRule) error
	CreateBatch(ctx context.Context, table *[]model.CasbinRule) error
	Delete(ctx context.Context, table *model.CasbinRule, args ...interface{}) error
	GetByColumns(ctx context.Context, tables *[]model.CasbinRule) error
}

type casbinRuleDao struct {
	db *gorm.DB
}

func NewCasbinRuleDao(db *gorm.DB) CasbinRuleDao {
	return &casbinRuleDao{db: db}
}

var _ CasbinRuleDao = (*casbinRuleDao)(nil)

func (d *casbinRuleDao) Create(ctx context.Context, table *model.CasbinRule) error {
	return d.db.WithContext(ctx).Create(table).Error
}
func (d *casbinRuleDao) CreateBatch(ctx context.Context, table *[]model.CasbinRule) error {
	return d.db.WithContext(ctx).Create(table).Error
}

func (d *casbinRuleDao) GetByColumns(ctx context.Context, tables *[]model.CasbinRule) error {
	return d.db.WithContext(ctx).Find(tables).Error
}

func (d *casbinRuleDao) Delete(ctx context.Context, table *model.CasbinRule, args ...interface{}) error {
	return d.db.WithContext(ctx).Delete(table, args...).Error
}
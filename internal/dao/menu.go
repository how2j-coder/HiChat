package dao

import (
	"com/chat/service/internal/model"
	"context"
	"gorm.io/gorm"
)

type MenuDao interface {
	Create(ctx context.Context, table *model.Menu) error
	UpdateByID(ctx context.Context, table *model.Menu) error
	DeleteByID(ctx context.Context, table *model.Menu) error
}

type menuDao struct {
	db *gorm.DB
}

var _ MenuDao = (*menuDao)(nil)

// NewMenuDao creating the dao interface
func NewMenuDao(db *gorm.DB) MenuDao {
	return &menuDao{
		db:    db,
	}
}

func (m *menuDao) Create(ctx context.Context, table *model.Menu) error  {
	return nil
}

func (m *menuDao) UpdateByID(ctx context.Context, table *model.Menu) error  {
	return nil
}

func (m *menuDao) DeleteByID(ctx context.Context, table *model.Menu) error  {
	return nil
}
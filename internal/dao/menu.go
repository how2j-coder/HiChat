package dao

import (
	"com/chat/service/internal/model"
	"context"
	"gorm.io/gorm"
)

type MenuDao interface {
	Create(ctx context.Context, table *model.Menu) error
	UpdateByID(ctx context.Context, table *model.Menu, update map[string]interface{}) error
	DeleteByID(ctx context.Context, table *model.Menu) error
	GetByMenuCode(ctx context.Context, code string) (*model.Menu, error)
	GetByParentID(ctx context.Context, parentID uint64) (*[]model.Menu, error)
	GetByMenuCodeExcID(ctx context.Context, id uint64, menuCode string) (*model.Menu, error)
}

type menuDao struct {
	db *gorm.DB
}

var _ MenuDao = (*menuDao)(nil)

// NewMenuDao creating the dao interface
func NewMenuDao(db *gorm.DB) MenuDao {
	return &menuDao{
		db: db,
	}
}

func (m *menuDao) Create(ctx context.Context, table *model.Menu) error {
	err := m.db.WithContext(ctx).Create(&table).Error
	return err
}

func (m *menuDao) GetByMenuCode(ctx context.Context, code string) (*model.Menu, error) {
	var menu = new(model.Menu)
	err := m.db.WithContext(ctx).First(menu, "menu_code = ?", code).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (m *menuDao) UpdateByID(ctx context.Context, table *model.Menu, update map[string]interface{}) error {
	err := m.db.WithContext(ctx).Model(table).Updates(update).Error
	return err
}

func (m *menuDao) DeleteByID(ctx context.Context, table *model.Menu) error {
	err := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.Menu{}, "id = ? or parent_menu_id = ?", table.ID, table.ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (m *menuDao) GetByParentID(ctx context.Context, parentID uint64) (*[]model.Menu, error) {
	var menus *[]model.Menu
	err := m.db.WithContext(ctx).Where("parent_menu_id = ?", parentID).Find(&menus).Error
	return menus, err
}

func (m *menuDao) GetByMenuCodeExcID(ctx context.Context, id uint64, menuCode string) (*model.Menu, error) {
	var menu = new(model.Menu)
	err := m.db.WithContext(ctx).First(menu).Where("id = ?", id).Where("menu_code = ?", menuCode).Error
	return menu, err
}

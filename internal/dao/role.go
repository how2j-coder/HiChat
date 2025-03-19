package dao

import (
	"com/chat/service/internal/model"
	"com/chat/service/pkg/datastore/query"
	"context"
	"errors"
	"gorm.io/gorm"
)

type RoleDao interface {
	Create(ctx context.Context, table *model.Role) error
	GetByName(ctx context.Context, name string) (*model.Role, error)
	GetByNameExcID(ctx context.Context, name string, id uint64) (*model.Role, error)
	UpdateByID(ctx context.Context, table *model.Role, update interface{}) error
	DeleteByID(ctx context.Context, id uint64) error
	GetByColumns(ctx context.Context, params *query.Params) (*[]model.Role, int64, error)
	SetUserRole(ctx context.Context, roleID []uint64, userID uint64) error
	SetMenuRole(ctx context.Context, roleID uint64, menuID []uint64) error
}

type roleDao struct {
	db *gorm.DB
}

var _ RoleDao = (*roleDao)(nil)

func NewRoleDao(db *gorm.DB) RoleDao {
	return &roleDao{
		db: db,
	}
}

func (d *roleDao) Create(ctx context.Context, table *model.Role) error {
	return d.db.WithContext(ctx).Create(&table).Error
}

func (d *roleDao) GetByName(ctx context.Context, name string) (*model.Role, error) {
	var role model.Role
	err := d.db.WithContext(ctx).First(&role, "role_name = ?", name).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (d *roleDao) UpdateByID(ctx context.Context, table *model.Role, update interface{}) error {
	return d.db.WithContext(ctx).Model(table).Updates(update).Error
}

func (d *roleDao) GetByNameExcID(ctx context.Context, name string, id uint64) (*model.Role, error) {
	var role model.Role
	err := d.db.WithContext(ctx).First(&role, "role_name = ?", name).Not("id <> ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (d *roleDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.Role{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (d *roleDao) GetByColumns(ctx context.Context, params *query.Params) (*[]model.Role, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}
	var total int64
	if params.Sort != "ignore count" {
		err = d.db.WithContext(ctx).Model(&model.Role{}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, errors.New("query count error: " + err.Error())
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	var roles *[]model.Role
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&roles).Error

	if err != nil {
		return nil, 0, errors.New("query role error: " + err.Error())
	}
	return roles, total, nil
}

func (d *roleDao) SetUserRole(ctx context.Context, roleID []uint64, userID uint64) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := d.db.WithContext(ctx).Delete(&[]model.RoleUser{}, "user_id = ?", userID).Error
		if err != nil {
			return err
		}
		var userRole []model.RoleUser
		for _, roleId := range roleID {
			userRole = append(userRole, model.RoleUser{RoleID: roleId, UserID: userID})
		}

		err = tx.Create(&userRole).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *roleDao) SetMenuRole(ctx context.Context, roleID uint64, menuID []uint64) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := d.db.WithContext(ctx).Delete(&[]model.RoleMenu{}, "role_id = ?", roleID).Error
		if err != nil {
			return err
		}
		var roleMenu []model.RoleMenu
		for _, menuId := range menuID {
			roleMenu = append(roleMenu, model.RoleMenu{MenuID: menuId, RoleID: roleID})
		}

		err = tx.Create(&roleMenu).Error
		if err != nil {
			return err
		}
		return nil
	})
}

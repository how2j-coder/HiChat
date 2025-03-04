package dao

import (
	"com/chat/service/internal/model"
	"context"
	"gorm.io/gorm"
)

type UserDao interface {
	Create(ctx context.Context, table *model.User) error
}

type userDao struct {
	db *gorm.DB
}

var _ UserDao = (*userDao)(nil)

// NewTeachDao creating the dao interface
func NewTeachDao(db *gorm.DB) UserDao {
	return &userDao{
		db:    db,
	}
}

// Create a record, insert the record, and the ID value is written back to the table.
func (d *userDao) Create(ctx context.Context, table *model.User) error {
	err := d.db.WithContext(ctx).Create(table).Error
	return err
}
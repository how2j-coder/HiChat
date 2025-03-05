package dao

import (
	"com/chat/service/internal/model"
	"context"
	"gorm.io/gorm"
)

type UserDao interface {
	Create(ctx context.Context, table *model.User) error
	UpdateByID(ctx context.Context, table *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByAccount(ctx context.Context, email string) (*model.User, error)
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

// UpdateByID the ID
func (d *userDao) UpdateByID(ctx context.Context, table *model.User) error {
	err := d.db.WithContext(ctx).Model(table).Updates(table).Error
	return err
}

// FindByEmail 查找用户
func (d *userDao) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).First(&user, "email = ?", email).Error
	return &user, err
}

// FindByAccount 查找用户
func (d *userDao) FindByAccount(ctx context.Context, account string) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).First(&user, "account = ?", account).Error
	return &user, err
}



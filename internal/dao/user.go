package dao

import (
	"com/chat/service/internal/model"
	"context"
)

type UserDao interface {
	Create(ctx context.Context, table *model.User) error
}

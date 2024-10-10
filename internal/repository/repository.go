package repository

import (
	"context"

	"github.com/greenblat17/auth/internal/model"
)

// UserRepository is a repository for User
type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Update(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id int64) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}

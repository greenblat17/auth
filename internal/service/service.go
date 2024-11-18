package service

import (
	"context"

	"github.com/greenblat17/auth/internal/model"
)

// UserService is a interface that provides method for User
type UserService interface {
	Create(ctx context.Context, userInfo *model.UserInfo) (int64, error)
	Update(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id int64) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}

// UserSaverProducer is an interface that provides method for User Saver Producer
type UserSaverProducer interface {
	Send(ctx context.Context, userInfo *model.UserInfo) error
}

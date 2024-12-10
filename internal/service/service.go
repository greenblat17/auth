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

// AuthService represents a service to authenticate user
type AuthService interface {
	Login(ctx context.Context, username, password string) (string, error)
	RefreshToken(ctx context.Context, oldRefreshToken string) (string, error)
	AccessToken(ctx context.Context, refreshToken string) (string, error)
}

// AccessService represents a service to check access
type AccessService interface {
	Check(ctx context.Context, accessToken, endpoint string) error
}

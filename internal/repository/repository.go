package repository

import (
	"context"
	"errors"
	"time"

	"github.com/greenblat17/auth/internal/model"
)

var (
	// ErrUserNotFound user not found
	ErrUserNotFound = errors.New("user not found")
	// ErrRuleNotFound rule not found
	ErrRuleNotFound = errors.New("rule not found")
)

// UserRepository is a repository for User
type UserRepository interface {
	Create(ctx context.Context, user *model.UserInfo) (int64, error)
	Update(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id int64) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}

// AuditRepository is a repository for Audit
type AuditRepository interface {
	Save(ctx context.Context, audit *model.Audit) error
}

// UserCacheRepository is a repository for caching user entity
type UserCacheRepository interface {
	Set(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	Expire(ctx context.Context, id int64, ttl time.Duration) error
}

// AccessRepository is a repository for access rules
type AccessRepository interface {
	GetByEndpoint(ctx context.Context, endpoint string) (*model.AccessRule, error)
}

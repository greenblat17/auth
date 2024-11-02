package redis

import (
	"context"
	"strconv"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/greenblat17/auth/internal/client/cache"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/repository"
	"github.com/greenblat17/auth/internal/repository/user/redis/converter"
	modelRepo "github.com/greenblat17/auth/internal/repository/user/redis/model"
)

type repo struct {
	cl cache.RedisClient
}

// NewRepository returns new user cache repo
func NewRepository(cl cache.RedisClient) repository.UserCacheRepository {
	return &repo{cl: cl}
}

func (r *repo) Set(ctx context.Context, user *model.User) (int64, error) {
	err := r.cl.HashSet(ctx, strconv.FormatInt(user.ID, 10), converter.ToUserFromService(user))
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	values, err := r.cl.HGetAll(ctx, strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, repository.ErrUserNotFound
	}

	var user modelRepo.User
	err = redigo.ScanStruct(values, &user)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	err := r.cl.Delete(ctx, strconv.FormatInt(id, 10))
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Expire(ctx context.Context, id int64, ttl time.Duration) error {
	err := r.cl.Expire(ctx, strconv.FormatInt(id, 10), ttl)
	if err != nil {
		return err
	}

	return nil
}

package user

import (
	"context"
	"errors"

	"github.com/greenblat17/auth/internal/model"
)

func (s *service) getUserFromCache(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userCacheRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return nil, errors.New("user not found in cache")
}

func (s *service) setUserToCache(ctx context.Context, user *model.User) error {
	id, err := s.userCacheRepository.Set(ctx, user)
	if err != nil {
		return err
	}

	err = s.userCacheRepository.Expire(ctx, id, s.ttl)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) deleteUserFromCache(ctx context.Context, id int64) error {
	err := s.userCacheRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

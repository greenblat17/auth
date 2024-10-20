package user

import (
	"context"

	"github.com/greenblat17/auth/internal/model"
)

func (s *service) Update(ctx context.Context, user *model.User) error {
	err := s.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

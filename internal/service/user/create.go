package user

import (
	"context"

	"github.com/greenblat17/auth/internal/model"
)

func (s *service) Create(ctx context.Context, userInfo *model.UserInfo) (int64, error) {
	id, err := s.userRepository.Create(ctx, userInfo)
	if err != nil {
		return 0, err
	}

	return id, nil
}

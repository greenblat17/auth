package auth

import (
	"context"
	"errors"

	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/pkg/auth"
)

func (s *service) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepository.Get(ctx, &model.UserFilter{Name: username})
	if err != nil {
		return "", err
	}

	isValid := auth.VerifyPassword(user.Info.Password, password)
	if !isValid {
		return "", errors.New("username or password is not valid")
	}

	refreshToken, err := auth.GenerateToken(user, s.refreshTokenConfig.Secret(), s.refreshTokenConfig.TTL())
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

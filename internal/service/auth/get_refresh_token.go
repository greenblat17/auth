package auth

import (
	"context"

	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/pkg/auth"
)

func (s *service) RefreshToken(_ context.Context, oldRefreshToken string) (string, error) {
	claims, err := auth.VerifyToken(oldRefreshToken, s.refreshTokenConfig.Secret())
	if err != nil {
		return "", err
	}

	user := &model.User{
		Info: model.UserInfo{
			Name: claims.Username,
			Role: claims.Role,
		},
	}

	refreshToken, err := auth.GenerateToken(user, s.refreshTokenConfig.Secret(), s.refreshTokenConfig.TTL())
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

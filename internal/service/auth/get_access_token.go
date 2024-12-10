package auth

import (
	"context"

	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/pkg/auth"
)

// AccessToken returnrs new access token
func (s *service) AccessToken(_ context.Context, refreshToken string) (string, error) {
	claims, err := auth.VerifyToken(refreshToken, s.refreshTokenConfig.Secret())
	if err != nil {
		return "", err
	}

	user := &model.User{
		Info: model.UserInfo{
			Name: claims.Username,
			Role: claims.Role,
		},
	}

	accessToken, err := auth.GenerateToken(user, s.accessTokenConfig.Secret(), s.accessTokenConfig.TTL())
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

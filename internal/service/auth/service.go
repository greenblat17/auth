package auth

import (
	"github.com/greenblat17/auth/internal/config"
	"github.com/greenblat17/auth/internal/repository"
	def "github.com/greenblat17/auth/internal/service"
)

var _ def.AuthService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository

	refreshTokenConfig config.TokenConfig
	accessTokenConfig  config.TokenConfig
}

// NewService returns a new auth service
func NewService(
	userRepository repository.UserRepository,
	refreshTokenConfig config.TokenConfig,
	accessTokenConfig config.TokenConfig,
) *service {
	return &service{
		userRepository:     userRepository,
		refreshTokenConfig: refreshTokenConfig,
		accessTokenConfig:  accessTokenConfig,
	}
}

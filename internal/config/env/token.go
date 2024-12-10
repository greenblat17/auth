package env

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/greenblat17/auth/internal/config"
)

const (
	refreshTokenTTLEnv    = "REFRESH_TOKEN_TTL"    //nolint:gosec
	refreshTokenSecretEnv = "REFRESH_TOKEN_SECRET" //nolint:gosec

	accessTokenTTLEnv    = "ACCESS_TOKEN_TTL"    //nolint:gosec
	accessTokenSecretEnv = "ACCESS_TOKEN_SECRET" //nolint:gosec
)

var _ config.TokenConfig = (*tokenConfig)(nil)

type tokenConfig struct {
	ttl    time.Duration
	secret []byte
}

// NewRefreshTokenConfig creates a new token configuration for refresh tokens
func NewRefreshTokenConfig() (*tokenConfig, error) {
	ttlStr, ok := os.LookupEnv(refreshTokenTTLEnv)
	if !ok {
		return nil, errors.New("refresh token ttl not found")
	}

	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		return nil, errors.New("failed to parse refresh token ttl")
	}

	secret, ok := os.LookupEnv(refreshTokenSecretEnv)
	if !ok {
		return nil, errors.New("refresh token secret not found")
	}

	return &tokenConfig{
		ttl:    time.Duration(ttl) * time.Minute,
		secret: []byte(secret),
	}, nil
}

// NewAccessTokenConfig creates a new token configuration for access tokens
func NewAccessTokenConfig() (*tokenConfig, error) {
	ttlStr, ok := os.LookupEnv(accessTokenTTLEnv)
	if !ok {
		return nil, errors.New("access token ttl not found")
	}

	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		return nil, errors.New("failed to parse access token ttl")
	}

	secret, ok := os.LookupEnv(accessTokenSecretEnv)
	if !ok {
		return nil, errors.New("access token secret not found")
	}

	return &tokenConfig{
		ttl:    time.Duration(ttl) * time.Minute,
		secret: []byte(secret),
	}, nil
}

// TTL indicates the period when token is valid
func (cfg *tokenConfig) TTL() time.Duration {
	return cfg.ttl
}

// Secret used to sign token and check its validity
func (cfg *tokenConfig) Secret() []byte {
	return cfg.secret
}

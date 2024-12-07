package access

import (
	"context"
	"errors"

	"github.com/greenblat17/auth/pkg/auth"
)

func (s *service) Check(ctx context.Context, accessToken, endpoint string) error {
	claims, err := auth.VerifyToken(accessToken, []byte("access secret key"))
	if err != nil {
		return err
	}

	accessRule, err := s.accessRepository.GetByEndpoint(ctx, endpoint)
	if err != nil {
		return err
	}

	if _, ok := accessRule.Role[claims.Role]; !ok {
		return errors.New("user doesn't have access")
	}

	return nil
}

package user

import (
	"context"

	"github.com/greenblat17/auth/internal/converter"
	"github.com/greenblat17/auth/internal/model"
)

func (s *service) Update(ctx context.Context, user *model.User) error {
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error

		errTx = s.userRepository.Update(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.auditRepository.Save(ctx, converter.ToAuditFromEntity(model.UserEntityType, "update"))
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

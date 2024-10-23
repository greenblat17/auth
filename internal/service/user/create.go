package user

import (
	"context"

	"github.com/greenblat17/auth/internal/converter"
	"github.com/greenblat17/auth/internal/model"
)

func (s *service) Create(ctx context.Context, userInfo *model.UserInfo) (int64, error) {
	var id int64

	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error

		id, errTx = s.userRepository.Create(ctx, userInfo)
		if errTx != nil {
			return errTx
		}

		errTx = s.auditRepository.Save(ctx, converter.ToAuditFromEntity(model.UserEntityType, "create"))
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

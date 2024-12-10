package user

import (
	"context"

	"github.com/greenblat17/auth/internal/converter"
	"github.com/greenblat17/auth/internal/model"
)

func (s *service) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.getUserFromCache(ctx, id)
	if nil == err {
		return user, nil
	}

	err = s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error

		user, errTx = s.userRepository.Get(ctx, &model.UserFilter{ID: id})
		if errTx != nil {
			return errTx
		}

		errTx = s.setUserToCache(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.auditRepository.Save(ctx, converter.ToAuditFromEntity(model.UserEntityType, "get"))
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

package tests

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/repository"
	"github.com/greenblat17/auth/internal/repository/mocks"
	"github.com/greenblat17/auth/internal/service"
	"github.com/greenblat17/auth/internal/service/user"
	"github.com/greenblat17/platform-common/pkg/db"
	dbMocks "github.com/greenblat17/platform-common/pkg/db/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		id  int64
	}

	type deps struct {
		userSaverProducer   service.UserSaverProducer
		userCacheRepository repository.UserCacheRepository
		userRepository      repository.UserRepository
		auditRepository     repository.AuditRepository
		txManager           db.TxManager
		ttl                 time.Duration
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id int64 = 1

		audit = &model.Audit{
			Entity: model.UserEntityType,
			Action: "delete",
		}

		ttl = time.Second
	)

	tests := []struct {
		name     string
		args     args
		mockFunc func(mc *minimock.Controller) deps
		err      error
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.DeleteMock.Expect(ctx, id).Return(nil)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)
				auditRepoMock.SaveMock.Expect(ctx, audit).Return(nil)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)
				userCacheRepoMock.DeleteMock.Expect(ctx, id).Return(nil)

				txManagerMock := dbMocks.NewTxManagerMock(mc)
				txManagerMock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				return deps{
					userCacheRepository: userCacheRepoMock,
					userRepository:      userRepoMock,
					auditRepository:     auditRepoMock,
					txManager:           txManagerMock,
					ttl:                 ttl,
				}
			},
			err: nil,
		},
		{
			name: "UserRepository return error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.DeleteMock.Expect(ctx, id).Return(assert.AnError)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)

				txManagerMock := dbMocks.NewTxManagerMock(mc)
				txManagerMock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				return deps{
					userCacheRepository: userCacheRepoMock,
					userRepository:      userRepoMock,
					auditRepository:     auditRepoMock,
					txManager:           txManagerMock,
					ttl:                 ttl,
				}
			},
			err: assert.AnError,
		},
		{
			name: "AuditRepository return error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.DeleteMock.Expect(ctx, id).Return(nil)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)
				auditRepoMock.SaveMock.Expect(ctx, audit).Return(assert.AnError)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)
				userCacheRepoMock.DeleteMock.Expect(ctx, id).Return(nil)

				txManagerMock := dbMocks.NewTxManagerMock(mc)
				txManagerMock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				return deps{
					userCacheRepository: userCacheRepoMock,
					userRepository:      userRepoMock,
					auditRepository:     auditRepoMock,
					txManager:           txManagerMock,
					ttl:                 ttl,
				}
			},
			err: assert.AnError,
		},
		{
			name: "UserCacheRepository return error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.DeleteMock.Expect(ctx, id).Return(nil)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)
				userCacheRepoMock.DeleteMock.Expect(ctx, id).Return(assert.AnError)

				txManagerMock := dbMocks.NewTxManagerMock(mc)
				txManagerMock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				return deps{
					userCacheRepository: userCacheRepoMock,
					userRepository:      userRepoMock,
					auditRepository:     auditRepoMock,
					txManager:           txManagerMock,
					ttl:                 ttl,
				}
			},
			err: assert.AnError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			deps := tt.mockFunc(mc)

			userSrv := user.NewService(
				deps.userSaverProducer,
				deps.userCacheRepository,
				deps.auditRepository,
				deps.userRepository,
				deps.txManager,
				deps.ttl,
			)

			err := userSrv.Delete(tt.args.ctx, tt.args.id)

			require.ErrorIs(t, err, tt.err)
		})
	}
}

package tests

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/repository"
	"github.com/greenblat17/auth/internal/repository/mocks"
	userService "github.com/greenblat17/auth/internal/service/user"
	"github.com/greenblat17/platform-common/pkg/db"
	dbMocks "github.com/greenblat17/platform-common/pkg/db/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		id  int64
	}

	type deps struct {
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

		user = &model.User{
			ID: id,
		}

		audit = &model.Audit{
			Entity: model.UserEntityType,
			Action: "get",
		}

		ttl = time.Second
	)

	tests := []struct {
		name     string
		args     args
		mockFunc func(mc *minimock.Controller) deps
		want     *model.User
		err      error
	}{
		{
			name: "User found in cache, Success",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)
				userCacheRepoMock.GetMock.Expect(ctx, id).Return(user, nil)

				txManagerMock := dbMocks.NewTxManagerMock(mc)

				return deps{
					userCacheRepository: userCacheRepoMock,
					userRepository:      userRepoMock,
					auditRepository:     auditRepoMock,
					txManager:           txManagerMock,
					ttl:                 ttl,
				}
			},
			want: user,
			err:  nil,
		},
		{
			name: "User not found in cache, Success",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.GetMock.Expect(ctx, id).Return(user, nil)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)
				userCacheRepoMock.GetMock.Expect(ctx, id).Return(nil, assert.AnError)
				userCacheRepoMock.SetMock.Expect(ctx, user).Return(id, nil)
				userCacheRepoMock.ExpireMock.Expect(ctx, id, ttl).Return(nil)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)
				auditRepoMock.SaveMock.Expect(ctx, audit).Return(nil)

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
			want: user,
			err:  nil,
		},
		{
			name: "UserRepository return error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.GetMock.Expect(ctx, id).Return(nil, assert.AnError)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)
				userCacheRepoMock.GetMock.Expect(ctx, id).Return(nil, assert.AnError)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)

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
			want: nil,
			err:  assert.AnError,
		},
		{
			name: "AuditRepository return error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.GetMock.Expect(ctx, id).Return(user, nil)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)
				userCacheRepoMock.GetMock.Expect(ctx, id).Return(nil, assert.AnError)
				userCacheRepoMock.SetMock.Expect(ctx, user).Return(id, nil)
				userCacheRepoMock.ExpireMock.Expect(ctx, id, ttl).Return(nil)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)
				auditRepoMock.SaveMock.Expect(ctx, audit).Return(assert.AnError)

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
			want: nil,
			err:  assert.AnError,
		},
		{
			name: "Set user to cache return error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.GetMock.Expect(ctx, id).Return(user, nil)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)
				userCacheRepoMock.GetMock.Expect(ctx, id).Return(nil, assert.AnError)
				userCacheRepoMock.SetMock.Expect(ctx, user).Return(int64(0), assert.AnError)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)

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
			want: nil,
			err:  assert.AnError,
		},
		{
			name: "set expire time in cache return error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.GetMock.Expect(ctx, id).Return(user, nil)

				userCacheRepoMock := mocks.NewUserCacheRepositoryMock(mc)
				userCacheRepoMock.GetMock.Expect(ctx, id).Return(nil, assert.AnError)
				userCacheRepoMock.SetMock.Expect(ctx, user).Return(id, nil)
				userCacheRepoMock.ExpireMock.Expect(ctx, id, ttl).Return(assert.AnError)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)

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
			want: nil,
			err:  assert.AnError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			deps := tt.mockFunc(mc)

			userSrv := userService.NewService(
				deps.userCacheRepository,
				deps.auditRepository,
				deps.userRepository,
				deps.txManager,
				deps.ttl,
			)

			got, err := userSrv.Get(tt.args.ctx, tt.args.id)

			if tt.err != nil {
				require.NotNil(t, err)
				assert.EqualError(t, err, tt.err.Error())
			} else {
				require.Equal(t, tt.want, got)
				assert.Nil(t, err)
			}
		})
	}
}

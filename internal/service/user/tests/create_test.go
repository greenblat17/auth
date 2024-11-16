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
	producerMocks "github.com/greenblat17/auth/internal/service/mocks"
	"github.com/greenblat17/auth/internal/service/user"
	"github.com/greenblat17/platform-common/pkg/db"
	dbMocks "github.com/greenblat17/platform-common/pkg/db/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx      context.Context
		userInfo *model.UserInfo
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
			Action: "create",
		}

		userInfo = &model.UserInfo{
			Name:     "testuser",
			Email:    "testuser@mail.com",
			Password: "testpass",
			Role:     model.RoleUser,
		}

		ttl = time.Second
	)

	tests := []struct {
		name     string
		args     args
		mockFunc func(mc *minimock.Controller) deps
		want     int64
		err      error
	}{
		{
			name: "Success",
			args: args{
				ctx:      ctx,
				userInfo: userInfo,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.CreateMock.Expect(ctx, userInfo).Return(id, nil)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)
				auditRepoMock.SaveMock.Expect(ctx, audit).Return(nil)

				txManagerMock := dbMocks.NewTxManagerMock(mc)
				txManagerMock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				userSaverProducerMock := producerMocks.NewUserSaverProducerMock(mc)
				userSaverProducerMock.SendMock.Expect(ctx, userInfo).Return(nil)

				return deps{
					userSaverProducer: userSaverProducerMock,
					userRepository:    userRepoMock,
					auditRepository:   auditRepoMock,
					txManager:         txManagerMock,
					ttl:               ttl,
				}
			},
			want: id,
			err:  nil,
		},
		{
			name: "UserRepository return error",
			args: args{
				ctx:      ctx,
				userInfo: userInfo,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.CreateMock.Expect(ctx, userInfo).Return(int64(0), assert.AnError)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)

				txManagerMock := dbMocks.NewTxManagerMock(mc)
				txManagerMock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				userSaverProducerMock := producerMocks.NewUserSaverProducerMock(mc)

				return deps{
					userSaverProducer: userSaverProducerMock,
					userRepository:    userRepoMock,
					auditRepository:   auditRepoMock,
					txManager:         txManagerMock,
					ttl:               ttl,
				}
			},
			want: 0,
			err:  assert.AnError,
		},
		{
			name: "AuditRepository return error",
			args: args{
				ctx:      ctx,
				userInfo: userInfo,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.CreateMock.Expect(ctx, userInfo).Return(id, nil)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)
				auditRepoMock.SaveMock.Expect(ctx, audit).Return(assert.AnError)

				txManagerMock := dbMocks.NewTxManagerMock(mc)
				txManagerMock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				userSaverProducerMock := producerMocks.NewUserSaverProducerMock(mc)

				return deps{
					userSaverProducer: userSaverProducerMock,
					userRepository:    userRepoMock,
					auditRepository:   auditRepoMock,
					txManager:         txManagerMock,
					ttl:               ttl,
				}
			},
			want: 0,
			err:  assert.AnError,
		},

		{
			name: "UserSaverProducer return error",
			args: args{
				ctx:      ctx,
				userInfo: userInfo,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userRepoMock := mocks.NewUserRepositoryMock(mc)
				userRepoMock.CreateMock.Expect(ctx, userInfo).Return(id, nil)

				auditRepoMock := mocks.NewAuditRepositoryMock(mc)
				auditRepoMock.SaveMock.Expect(ctx, audit).Return(nil)

				txManagerMock := dbMocks.NewTxManagerMock(mc)
				txManagerMock.ReadCommitedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})

				userSaverProducerMock := producerMocks.NewUserSaverProducerMock(mc)
				userSaverProducerMock.SendMock.Expect(ctx, userInfo).Return(assert.AnError)

				return deps{
					userSaverProducer: userSaverProducerMock,
					userRepository:    userRepoMock,
					auditRepository:   auditRepoMock,
					txManager:         txManagerMock,
					ttl:               ttl,
				}
			},
			want: 0,
			err:  assert.AnError,
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

			id, err := userSrv.Create(tt.args.ctx, tt.args.userInfo)

			if tt.err != nil {
				require.NotNil(t, err)

			} else {
				require.Equal(t, tt.want, id)
				require.Nil(t, err)
			}
		})
	}

}

package tests

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	userAPI "github.com/greenblat17/auth/internal/api/user"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/service"
	"github.com/greenblat17/auth/internal/service/mocks"
	desc "github.com/greenblat17/auth/pkg/user_v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	type deps struct {
		userService service.UserService
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       int64 = 1
		name           = "test"
		email          = "test@example.com"
		password       = "password"
		now            = time.Now()

		req = &desc.GetRequest{
			Id: id,
		}

		user = &model.User{
			ID: id,
			Info: model.UserInfo{
				Name:     name,
				Email:    email,
				Password: password,
				Role:     "USER",
			},
			CreatedAt: now,
			UpdatedAt: &now,
		}

		resp = &desc.GetResponse{
			User: &desc.User{
				Id: id,
				Info: &desc.UserInfo{
					Name:  name,
					Email: email,
					Role:  desc.Role_USER,
				},
				CreatedAt: timestamppb.New(now),
				UpdatedAt: timestamppb.New(now),
			},
		}
	)

	tests := []struct {
		name     string
		args     args
		mockFunc func(mc *minimock.Controller) deps
		want     *desc.GetResponse
		err      error
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				req: req,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userServiceMock := mocks.NewUserServiceMock(mc)
				userServiceMock.GetMock.Expect(ctx, id).Return(user, nil)

				return deps{
					userService: userServiceMock,
				}
			},
			want: resp,
			err:  nil,
		},
		{
			name: "UserService return error",
			args: args{
				ctx: ctx,
				req: req,
			},
			mockFunc: func(mc *minimock.Controller) deps {
				userServiceMock := mocks.NewUserServiceMock(mc)
				userServiceMock.GetMock.Expect(ctx, id).Return(nil, assert.AnError)

				return deps{
					userService: userServiceMock,
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

			userHandler := userAPI.NewImplementation(deps.userService)

			got, err := userHandler.Get(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

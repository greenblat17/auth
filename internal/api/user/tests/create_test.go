package tests

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/greenblat17/auth/internal/api/user"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/service"
	"github.com/greenblat17/auth/internal/service/mocks"
	desc "github.com/greenblat17/auth/pkg/user_v1"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	type deps struct {
		userService service.UserService
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       int64 = 1
		name           = "test_create"
		email          = "test@example.com"
		password       = "password"
		role           = model.RoleUser

		req = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.Role_USER,
		}

		userInfo = &model.UserInfo{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}

		resp = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name     string
		args     args
		mockFunc func(mc *minimock.Controller) deps
		want     *desc.CreateResponse
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
				userServiceMock.CreateMock.Expect(ctx, userInfo).Return(id, nil)

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
				userServiceMock.CreateMock.Expect(ctx, userInfo).Return(int64(0), assert.AnError)

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

			userAPI := user.NewImplementation(deps.userService)

			got, err := userAPI.Create(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

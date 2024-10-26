package tests

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	userAPI "github.com/greenblat17/auth/internal/api/user"
	"github.com/greenblat17/auth/internal/model"
	"github.com/greenblat17/auth/internal/service"
	"github.com/greenblat17/auth/internal/service/mocks"
	desc "github.com/greenblat17/auth/pkg/user_v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	type deps struct {
		userService service.UserService
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		name  = "test"
		email = "test@example.com"
		role  = "USER"

		user = &model.User{
			ID: 1,
			Info: model.UserInfo{
				Name:  name,
				Email: email,
				Role:  role,
			},
		}

		req = &desc.UpdateRequest{
			Id:    1,
			Name:  &wrapperspb.StringValue{Value: "test"},
			Email: &wrapperspb.StringValue{Value: "test@example.com"},
			Role:  desc.Role_USER,
		}
	)

	tests := []struct {
		name     string
		args     args
		mockFunc func(mc *minimock.Controller) deps
		want     *emptypb.Empty
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
				userServiceMock.UpdateMock.Expect(ctx, user).Return(nil)

				return deps{
					userService: userServiceMock,
				}
			},
			want: &emptypb.Empty{},
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
				userServiceMock.UpdateMock.Expect(ctx, user).Return(assert.AnError)

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

			api := userAPI.NewImplementation(deps.userService)

			got, err := api.Update(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

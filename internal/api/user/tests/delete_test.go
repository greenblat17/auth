package tests

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/greenblat17/auth/internal/api/user"
	"github.com/greenblat17/auth/internal/service"
	"github.com/greenblat17/auth/internal/service/mocks"
	desc "github.com/greenblat17/auth/pkg/user_v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	type deps struct {
		userService service.UserService
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id int64 = 1

		req = &desc.DeleteRequest{
			Id: id,
		}

		resp = &emptypb.Empty{}
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
				userServiceMock.DeleteMock.Expect(ctx, id).Return(nil)

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
				userServiceMock.DeleteMock.Expect(ctx, id).Return(assert.AnError)

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

			got, err := userAPI.Delete(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

package user

import (
	"context"

	"github.com/greenblat17/auth/internal/converter"
	desc "github.com/greenblat17/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update updates an existing user
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	err = i.userService.Update(ctx, converter.ToUserFromUpdateAPI(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

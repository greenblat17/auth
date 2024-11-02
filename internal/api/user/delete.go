package user

import (
	"context"

	desc "github.com/greenblat17/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a user by ID
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	err = i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

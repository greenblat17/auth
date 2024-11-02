package user

import (
	"context"

	"github.com/greenblat17/auth/internal/converter"
	desc "github.com/greenblat17/auth/pkg/user_v1"
)

// Get returns the user
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	userObj, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: converter.ToUserFromService(userObj),
	}, nil
}

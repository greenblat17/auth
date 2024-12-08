package user

import (
	"context"

	"github.com/greenblat17/auth/internal/converter"
	desc "github.com/greenblat17/auth/pkg/user_v1"
)

// Create creates a new user
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserInfoFromCreateAPI(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

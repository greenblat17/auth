package user

import (
	"github.com/greenblat17/auth/internal/service"
	desc "github.com/greenblat17/auth/pkg/user_v1"
)

// Implementation represents user server
type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// NewImplementation returns a new implementation of server
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}

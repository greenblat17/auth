package user

import (
	"github.com/greenblat17/auth/internal/repository"
	def "github.com/greenblat17/auth/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

// NewService returns a new service
func NewService(userRepository repository.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

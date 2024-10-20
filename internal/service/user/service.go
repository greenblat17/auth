package user

import (
	"github.com/greenblat17/auth/internal/client/db"
	"github.com/greenblat17/auth/internal/repository"
	def "github.com/greenblat17/auth/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

// NewService returns a new service
func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) *service {
	return &service{
		userRepository: userRepository,
		txManager:      txManager,
	}
}

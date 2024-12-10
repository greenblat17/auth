package access

import (
	"github.com/greenblat17/auth/internal/repository"
	def "github.com/greenblat17/auth/internal/service"
)

var _ def.AccessService = (*service)(nil)

type service struct {
	accessRepository repository.AccessRepository
}

// NewService returns a new access service
func NewService(accessRepository repository.AccessRepository) *service {
	return &service{accessRepository: accessRepository}
}

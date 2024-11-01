package user

import (
	"time"

	"github.com/greenblat17/auth/internal/repository"
	def "github.com/greenblat17/auth/internal/service"
	"github.com/greenblat17/platform-common/pkg/db"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userCacheRepository repository.UserCacheRepository
	auditRepository     repository.AuditRepository
	userRepository      repository.UserRepository
	txManager           db.TxManager
	ttl                 time.Duration
}

// NewService returns a new service
func NewService(
	userCacheRepository repository.UserCacheRepository,
	auditRepository repository.AuditRepository,
	userRepository repository.UserRepository,
	txManager db.TxManager,
	ttl time.Duration,
) *service {
	return &service{
		userCacheRepository: userCacheRepository,
		auditRepository:     auditRepository,
		userRepository:      userRepository,
		txManager:           txManager,
		ttl:                 ttl,
	}
}

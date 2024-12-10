package access

import (
	"github.com/greenblat17/auth/internal/service"
	desc "github.com/greenblat17/auth/pkg/access_v1"
)

// Implementation access handler
type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

// NewImplementation create access handler
func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}

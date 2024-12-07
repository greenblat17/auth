package auth

import (
	"github.com/greenblat17/auth/internal/service"
	desc "github.com/greenblat17/auth/pkg/auth_v1"
)

// Implementation auth hand;ler
type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

// NewImplementation returns a new auth handler
func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}
}

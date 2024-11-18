package env

import (
	"errors"
	"net"
	"os"

	"github.com/greenblat17/auth/internal/config"
)

const (
	swaggerHostEnvName = "SWAGGER_HOST"
	swaggerPortEnvName = "SWAGGER_PORT"
)

var _ config.SwaggerConfig = (*swaggerConfig)(nil)

type swaggerConfig struct {
	host string
	port string
}

// NewSwaggerConfig returns a new swagger config
func NewSwaggerConfig() (*swaggerConfig, error) {
	host, ok := os.LookupEnv(swaggerHostEnvName)
	if !ok {
		return nil, errors.New("swagger host not found")
	}

	port, ok := os.LookupEnv(swaggerPortEnvName)
	if !ok {
		return nil, errors.New("swagger port not found")
	}

	return &swaggerConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *swaggerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

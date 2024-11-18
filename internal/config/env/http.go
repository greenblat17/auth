package env

import (
	"errors"
	"net"
	"os"

	"github.com/greenblat17/auth/internal/config"
)

var _ config.HTTPConfig = (*httpConfig)(nil)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	host string
	port string
}

// NewHTTPConfig returns new http config
func NewHTTPConfig() (*httpConfig, error) {
	host, ok := os.LookupEnv(httpHostEnvName)
	if !ok {
		return nil, errors.New("http host not found")
	}

	port, ok := os.LookupEnv(httpPortEnvName)
	if !ok {
		return nil, errors.New("http port not found")
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

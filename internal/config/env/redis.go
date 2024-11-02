package env

import (
	"net"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	redisHostEnvName              = "REDIS_HOST"
	redisPortEnvName              = "REDIS_PORT"
	redisConnectionTimeoutEnvName = "REDIS_CONNECTION_TIMEOUT_SEC"
	redisMaxIdleEnvName           = "REDIS_MAX_IDLE"
	redisIdleTimeoutEnvName       = "REDIS_IDLE_TIMEOUT_SEC"
	redisTTLEnvName               = "REDIS_TTL_SEC"
)

type redisConfig struct {
	host string
	port string

	connectionTimeout time.Duration

	maxIdle     int
	idleTimeout time.Duration

	ttl time.Duration
}

// NewRedisConfig returns new redis configuration
func NewRedisConfig() (*redisConfig, error) {
	host, ok := os.LookupEnv(redisHostEnvName)
	if !ok {
		return nil, errors.New("redis host not found")
	}

	port, ok := os.LookupEnv(redisPortEnvName)
	if !ok {
		return nil, errors.New("redis port not found")
	}

	connectionTimeoutStr, ok := os.LookupEnv(redisConnectionTimeoutEnvName)
	if !ok {
		return nil, errors.New("redis connection timeout not found")
	}

	connectionTimeout, err := strconv.ParseInt(connectionTimeoutStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse connection timeout")
	}

	maxIdleStr, ok := os.LookupEnv(redisMaxIdleEnvName)
	if !ok {
		return nil, errors.New("redis max idle not found")
	}

	maxIdle, err := strconv.Atoi(maxIdleStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse max idle")
	}

	idleTimeoutStr, ok := os.LookupEnv(redisIdleTimeoutEnvName)
	if !ok {
		return nil, errors.New("redis idle timeout not found")
	}

	idleTimeout, err := strconv.ParseInt(idleTimeoutStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse idle timeout")
	}

	ttlStr, ok := os.LookupEnv(redisTTLEnvName)
	if !ok {
		return nil, errors.New("redis ttl not found")
	}

	ttl, err := strconv.ParseInt(ttlStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse ttl")
	}

	return &redisConfig{
		host:              host,
		port:              port,
		connectionTimeout: time.Duration(connectionTimeout) * time.Second,
		maxIdle:           maxIdle,
		idleTimeout:       time.Duration(idleTimeout) * time.Second,
		ttl:               time.Duration(ttl) * time.Second,
	}, nil
}

func (cfg *redisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *redisConfig) ConnectionTimeout() time.Duration {
	return cfg.connectionTimeout
}

func (cfg *redisConfig) MaxIdle() int {
	return cfg.maxIdle
}

func (cfg *redisConfig) IdleTimeout() time.Duration {
	return cfg.idleTimeout
}

func (cfg *redisConfig) TTL() time.Duration {
	return cfg.ttl
}

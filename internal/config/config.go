package config

import (
	"time"

	"github.com/joho/godotenv"
)

// GRPCConfig is the configuration for the gRPC server
type GRPCConfig interface {
	Address() string
}

// HTTPConfig is the configuration for HTTP server
type HTTPConfig interface {
	Address() string
}

// PGConfig is the configuration for the PostgreSQL database
type PGConfig interface {
	DSN() string
}

// RedisConfig is the configuration for Redis database
type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
	TTL() time.Duration
}

// KafkaConfig is the configuration for Kafka
type KafkaConfig interface {
	Brokers() []string
	Retry() int
}

// SwaggerConfig is the configuration for Swagger server
type SwaggerConfig interface {
	Address() string
}

// TokenConfig is the configuration for token generation
type TokenConfig interface {
	Secret() []byte
	TTL() time.Duration
}

// Load loads configuration from environment variables file
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

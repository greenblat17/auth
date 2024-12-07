package app

import (
	"log"

	"github.com/greenblat17/auth/internal/config"
	"github.com/greenblat17/auth/internal/config/env"
)

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %v", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %v", err)
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) ProducerConfig() config.KafkaConfig {
	if s.kafkaConfig == nil {
		cfg, err := env.NewKafkaConfig()
		if err != nil {
			log.Fatalf("failed to get producer config: %v", err)
		}

		s.kafkaConfig = cfg
	}

	return s.kafkaConfig
}

func (s *serviceProvider) RefreshTokenConfig() config.TokenConfig {
	if s.refreshTokenConfig == nil {
		cfg, err := env.NewRefreshTokenConfig()
		if err != nil {
			log.Fatalf("failed to get refresh token config: %v", err)
		}

		s.refreshTokenConfig = cfg
	}

	return s.refreshTokenConfig
}

func (s *serviceProvider) AccessTokenConfig() config.TokenConfig {
	if s.accessTokenConfig == nil {
		cfg, err := env.NewAccessTokenConfig()
		if err != nil {
			log.Fatalf("failed to get access token config: %v", err)
		}

		s.accessTokenConfig = cfg
	}

	return s.accessTokenConfig
}

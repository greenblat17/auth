package app

import (
	"context"
	"log"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/greenblat17/auth/internal/api/user"
	"github.com/greenblat17/auth/internal/client/cache"
	"github.com/greenblat17/auth/internal/client/cache/redis"
	"github.com/greenblat17/auth/internal/config"
	"github.com/greenblat17/auth/internal/config/env"
	"github.com/greenblat17/auth/internal/repository"
	"github.com/greenblat17/auth/internal/repository/audit"
	userRepository "github.com/greenblat17/auth/internal/repository/user/pg"
	userCacheRepository "github.com/greenblat17/auth/internal/repository/user/redis"
	"github.com/greenblat17/auth/internal/service"
	userService "github.com/greenblat17/auth/internal/service/user"
	"github.com/greenblat17/platform-common/pkg/closer"
	"github.com/greenblat17/platform-common/pkg/db"
	"github.com/greenblat17/platform-common/pkg/db/pg"
	"github.com/greenblat17/platform-common/pkg/db/transaction"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	grpcConfig  config.GRPCConfig
	redisConfig config.RedisConfig
	httpConfig  config.HTTPConfig

	dbClient  db.Client
	txManager db.TxManager

	redisPool   *redigo.Pool
	redisClient cache.RedisClient

	userCacheRepository repository.UserCacheRepository
	userRepository      repository.UserRepository
	auditRepository     repository.AuditRepository

	userService service.UserService

	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

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

func (s *serviceProvider) RedisPool() *redigo.Pool {
	if s.redisPool == nil {
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.RedisConfig().MaxIdle(),
			IdleTimeout: s.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", s.RedisConfig().Address())
			},
		}
	}

	return s.redisPool
}

func (s *serviceProvider) RedisClient() cache.RedisClient {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(s.RedisPool(), s.RedisConfig())
	}

	return s.redisClient
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create database client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) AuditRepository(ctx context.Context) repository.AuditRepository {
	if s.auditRepository == nil {
		s.auditRepository = audit.NewRepository(s.DBClient(ctx))
	}

	return s.auditRepository
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserCacheRepository() repository.UserCacheRepository {
	if s.userCacheRepository == nil {
		s.userCacheRepository = userCacheRepository.NewRepository(s.RedisClient())
	}

	return s.userCacheRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserCacheRepository(),
			s.AuditRepository(ctx),
			s.UserRepository(ctx),
			s.TxManager(ctx),
			s.RedisConfig().TTL(),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

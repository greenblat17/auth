package app

import (
	"context"
	"log"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/greenblat17/auth/internal/api/access"
	"github.com/greenblat17/auth/internal/api/auth"
	"github.com/greenblat17/auth/internal/api/user"
	"github.com/greenblat17/auth/internal/client/cache"
	"github.com/greenblat17/auth/internal/client/cache/redis"
	"github.com/greenblat17/auth/internal/client/kafka"
	"github.com/greenblat17/auth/internal/client/kafka/producer"
	"github.com/greenblat17/auth/internal/config"
	"github.com/greenblat17/auth/internal/repository"
	accessRepository "github.com/greenblat17/auth/internal/repository/access"
	"github.com/greenblat17/auth/internal/repository/audit"
	userRepository "github.com/greenblat17/auth/internal/repository/user/pg"
	userCacheRepository "github.com/greenblat17/auth/internal/repository/user/redis"
	"github.com/greenblat17/auth/internal/service"
	accessService "github.com/greenblat17/auth/internal/service/access"
	authService "github.com/greenblat17/auth/internal/service/auth"
	"github.com/greenblat17/auth/internal/service/producer/user_saver"
	userService "github.com/greenblat17/auth/internal/service/user"
	"github.com/greenblat17/platform-common/pkg/closer"
	"github.com/greenblat17/platform-common/pkg/db"
	"github.com/greenblat17/platform-common/pkg/db/pg"
	"github.com/greenblat17/platform-common/pkg/db/transaction"
)

const (
	userSaverProducerTopic = "user-save-topic"
)

type serviceProvider struct {
	pgConfig           config.PGConfig
	grpcConfig         config.GRPCConfig
	redisConfig        config.RedisConfig
	httpConfig         config.HTTPConfig
	swaggerConfig      config.SwaggerConfig
	kafkaConfig        config.KafkaConfig
	accessTokenConfig  config.TokenConfig
	refreshTokenConfig config.TokenConfig

	dbClient  db.Client
	txManager db.TxManager

	redisPool   *redigo.Pool
	redisClient cache.RedisClient

	producer          kafka.Producer
	userSaverProducer service.UserSaverProducer

	userCacheRepository repository.UserCacheRepository
	userRepository      repository.UserRepository
	auditRepository     repository.AuditRepository
	accessRepository    repository.AccessRepository

	userService   service.UserService
	authService   service.AuthService
	accessService service.AccessService

	userImpl   *user.Implementation
	authImpl   *auth.Implementation
	accessImpl *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
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

func (s *serviceProvider) Producer() kafka.Producer {
	if s.producer == nil {
		pr, err := producer.NewProducer(s.ProducerConfig())
		if err != nil {
			log.Fatalf("failed to create producer: %v", err)
		}

		closer.Add(pr.Close)

		s.producer = pr
	}

	return s.producer
}

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepository(s.DBClient(ctx))
	}

	return s.accessRepository
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

func (s *serviceProvider) UserSaverProducer() service.UserSaverProducer {
	if s.userSaverProducer == nil {
		s.userSaverProducer = user_saver.NewProducer(s.Producer(), userSaverProducerTopic)
	}

	return s.userSaverProducer
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserSaverProducer(),
			s.UserCacheRepository(),
			s.AuditRepository(ctx),
			s.UserRepository(ctx),
			s.TxManager(ctx),
			s.RedisConfig().TTL(),
		)
	}

	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.UserRepository(ctx), s.RefreshTokenConfig(), s.AccessTokenConfig())
	}

	return s.authService
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(s.AccessRepository(ctx))
	}

	return s.accessService
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AuthImplementation(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) AccessImplementation(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}

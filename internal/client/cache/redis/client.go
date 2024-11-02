package redis

import (
	"context"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/greenblat17/auth/internal/client/cache"
	"github.com/greenblat17/auth/internal/config"
)

const (
	hSetCommand    = "HSET"
	setCommand     = "SET"
	hGetAllCommand = "HGETALL"
	getCommand     = "GET"
	expireCommand  = "EXPIRE"
	pingCommand    = "PING"
	deleteCommand  = "DEL"
)

var _ cache.RedisClient = (*client)(nil)

type handler func(ctx context.Context, conn redis.Conn) error

type client struct {
	pool   *redis.Pool
	config config.RedisConfig
}

// NewClient returns new redis client with pool
func NewClient(pool *redis.Pool, config config.RedisConfig) *client {
	return &client{
		pool:   pool,
		config: config,
	}
}

func (c *client) HashSet(ctx context.Context, key string, values interface{}) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do(hSetCommand, redis.Args{key}.AddFlat(values)...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) Set(ctx context.Context, key string, value interface{}) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do(setCommand, redis.Args{key}.Add(value)...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) HGetAll(ctx context.Context, key string) ([]interface{}, error) {
	var value []interface{}
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		var errEx error
		value, errEx = redis.Values(conn.Do(hGetAllCommand, key))
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *client) Get(ctx context.Context, key string) (interface{}, error) {
	var value interface{}
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		var errEx error
		value, errEx = conn.Do(getCommand, key)
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do(expireCommand, key, int(expiration.Seconds()))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) Ping(ctx context.Context) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do(pingCommand)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) Delete(ctx context.Context, key string) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do(deleteCommand, key)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) execute(ctx context.Context, handler handler) error {
	conn, err := c.connect(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	err = handler(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) connect(ctx context.Context) (redis.Conn, error) {
	ctxConnTimeout, cancel := context.WithTimeout(ctx, c.config.ConnectionTimeout())
	defer cancel()

	conn, err := c.pool.GetContext(ctxConnTimeout)
	if err != nil {
		log.Printf("failed to get redis connection: %v", err)

		_ = conn.Close()
		return nil, err
	}

	return conn, nil
}

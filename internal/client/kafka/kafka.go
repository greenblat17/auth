package kafka

import (
	"context"

	"github.com/greenblat17/auth/internal/client/kafka/consumer"
)

// Consumer represents kafka consumer
type Consumer interface {
	Consume(ctx context.Context, topicName string, handler consumer.Handler) (err error)
	Close() error
}

// Producer represents kafka producer
type Producer interface {
	Produce(ctx context.Context, data []byte, topicName string) error
	Close() error
}

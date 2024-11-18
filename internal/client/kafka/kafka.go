package kafka

import (
	"context"
)

// Producer represents kafka producer
type Producer interface {
	Produce(ctx context.Context, data []byte, topicName string) error
	Close() error
}

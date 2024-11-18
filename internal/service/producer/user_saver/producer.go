package user_saver

import (
	"context"
	"encoding/json"

	"github.com/greenblat17/auth/internal/client/kafka"
	"github.com/greenblat17/auth/internal/model"
)

type service struct {
	producer  kafka.Producer
	topicName string
}

// NewProducer return new kafka user saver producer
func NewProducer(producer kafka.Producer, topicName string) *service {
	return &service{
		producer:  producer,
		topicName: topicName,
	}
}

func (s *service) Send(ctx context.Context, userInfo *model.UserInfo) error {
	data, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}

	err = s.producer.Produce(ctx, data, s.topicName)
	if err != nil {
		return err
	}

	return nil
}

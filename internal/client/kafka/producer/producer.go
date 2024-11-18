package producer

import (
	"context"
	"log"

	"github.com/IBM/sarama"
	"github.com/greenblat17/auth/internal/config"
)

type producer struct {
	producer sarama.SyncProducer
	config   config.KafkaConfig
}

// NewProducer returns a new kafka producer
func NewProducer(cfg config.KafkaConfig) (*producer, error) {
	syncProducer, err := newSyncProducer(cfg.Brokers(), cfg.Retry())
	if err != nil {
		return nil, err
	}

	return &producer{
		producer: syncProducer,
		config:   cfg,
	}, nil
}

func (p *producer) Produce(_ context.Context, data []byte, topicName string) error {
	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder(data),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("failed to send msg: %v\n", err)

		return err
	}

	log.Printf("msg sent to topic %s in partition %d with offset %d\n", msg.Topic, partition, offset)

	return nil
}

func (p *producer) Close() error {
	return p.producer.Close()
}

func newSyncProducer(brokerList []string, retryMax int) (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(brokerList, producerConfig(retryMax))
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func producerConfig(retryMax int) *sarama.Config {
	cfg := sarama.NewConfig()

	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = retryMax
	cfg.Producer.Return.Successes = true

	return cfg
}

package env

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	brokersAddressEnvName = "KAFKA_BROKERS_ADDRESS"
	retryMaxEnvName       = "KAFKA_RETRY_MAX"
)

type kafkaConfig struct {
	brokers  []string
	retryMax int
}

// NewKafkaConfig return configuration for Kafka
func NewKafkaConfig() (*kafkaConfig, error) {
	brokerAddress, ok := os.LookupEnv(brokersAddressEnvName)
	if !ok {
		return nil, errors.New("kafka brokers address not found")
	}

	retryMaxStr, ok := os.LookupEnv(retryMaxEnvName)
	if !ok {
		return nil, errors.New("kafka retry max count not found")
	}
	retryMax, err := strconv.Atoi(retryMaxStr)
	if err != nil {
		return nil, errors.New("failed to parse kafka retry max count")
	}

	return &kafkaConfig{
		brokers:  strings.Split(brokerAddress, ","),
		retryMax: retryMax,
	}, nil
}

func (cfg *kafkaConfig) Brokers() []string {
	return cfg.brokers
}

func (cfg *kafkaConfig) Retry() int {
	return cfg.retryMax
}

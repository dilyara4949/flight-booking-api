package kafka_client

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dilyara4949/flight-booking-api/internal/config"
)

func ConnectProducer(cfg config.Kafka) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	return p, nil
}

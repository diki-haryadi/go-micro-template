package kafkaHealthCheckUseCase

import (
	"github.com/segmentio/kafka-go"

	healthCheckDomain "github.com/diki-haryadi/go-micro-template/internal/health_check/domain"
	"github.com/diki-haryadi/ztools/config"
)

type useCase struct{}

func NewUseCase() healthCheckDomain.KafkaHealthCheckUseCase {
	return &useCase{}
}

func (uc *useCase) Check() bool {
	brokers := kafka.TCP(config.BaseConfig.Kafka.ClientBrokers...)

	conn, err := kafka.Dial(brokers.Network(), brokers.String())
	if err != nil {
		return false
	}

	_ = conn.Close()

	return true
}

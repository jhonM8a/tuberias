package factory

import (
	"errors"

	rabbitMq "tuberias/infraestructure/brokers"
	interfaceInfraestucture "tuberias/infraestructure/interface"
)

const (
	brokerRabbitMQ = "rabbitmq"
)

func NewBroker(brokerType string, config string) (interfaceInfraestucture.Broker, error) {
	switch brokerType {
	case brokerRabbitMQ:
		return rabbitMq.NewRabbitMQBroker(config), nil
	default:
		return nil, errors.New("unsupported broker type")
	}
}

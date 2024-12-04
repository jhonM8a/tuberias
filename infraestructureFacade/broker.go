package infraestructurefacade

import (
	"errors"

	rabbitMq "tuberias/infraestructure"
)

type Broker interface {
	Connect() error
	Publish(queue string, message []byte) error
	Consume(queue string) (<-chan []byte, error)
	Close() error
}

func NewBroker(brokerType string, config string) (Broker, error) {
	switch brokerType {
	case "rabbitmq":
		return rabbitMq.NewRabbitMQBroker(config), nil
	default:
		return nil, errors.New("unsupported broker type")
	}
}

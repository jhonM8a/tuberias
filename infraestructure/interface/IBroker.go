package interfaceInfraestucture

type Broker interface {
	Connect() error
	Publish(queue string, message []byte) error
	Consume(queue string) (<-chan []byte, error)
	Close() error
}

package brokers

import (
	"github.com/streadway/amqp"
)

type RabbitMQBroker struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	url        string
}

func NewRabbitMQBroker(url string) *RabbitMQBroker {
	return &RabbitMQBroker{url: url}
}

func (r *RabbitMQBroker) Connect() error {
	conn, err := amqp.Dial(r.url)
	if err != nil {
		return err
	}
	r.connection = conn

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	r.channel = ch

	return nil
}

func (r *RabbitMQBroker) Publish(queue string, message []byte) error {
	_, err := r.channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}

func (r *RabbitMQBroker) Consume(queue string) (<-chan []byte, error) {
	msgs, err := r.channel.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return nil, err
	}

	out := make(chan []byte)
	go func() {
		for d := range msgs {
			out <- d.Body
		}
	}()

	return out, nil
}

func (r *RabbitMQBroker) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	return r.connection.Close()
}

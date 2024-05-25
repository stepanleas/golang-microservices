package rabbit_mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqClient struct {
	conn *amqp.Connection
}

func NewRabbitMqClient(url string) (*RabbitMqClient, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return &RabbitMqClient{
		conn: conn,
	}, nil
}

func (r *RabbitMqClient) Close() {
	r.conn.Close()
}

func (r *RabbitMqClient) Channel() (*amqp.Channel, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

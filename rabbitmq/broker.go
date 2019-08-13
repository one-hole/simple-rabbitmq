package rabbitmq

import (
	"github/one-hole/simple-rabbitmq/brokers"

	"github.com/streadway/amqp"
)

type rabbitBroker struct {
	conn *amqp.Connection
}

func Dial(url string) (brokers.Broker, error) {
	conn, err := amqp.Dial(url)

	if err != nil {
		return nil, err
	}

	return &rabbitBroker{conn: conn}, nil
}

func (b *rabbitBroker) Close() {
	_ = b.conn.Close()
}

func (b *rabbitBroker) Subscribe(topic string, routingKey string, handler brokers.MessageHandler) error {
	return nil
}

//Subscribe(topic, identity string, handler MessageHandler)

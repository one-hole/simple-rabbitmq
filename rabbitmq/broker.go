package rabbitmq

import (
	"github/one-hole/simple-rabbitmq/brokers"

	"github.com/streadway/amqp"
)

type broker struct {
	conn *amqp.Connection
}

func Dial() (brokers.Broker, error) {
	return nil, nil
}

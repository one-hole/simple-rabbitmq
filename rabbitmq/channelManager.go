package rabbitmq

import "github.com/streadway/amqp"

type channelManager struct {
	conn     *amqp.Connection
	channels []*amqp.Channel
}

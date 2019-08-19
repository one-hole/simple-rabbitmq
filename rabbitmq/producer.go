package rabbitmq

import (
	"github.com/one-hole/simple-rabbitmq/brokers"

	"github.com/streadway/amqp"
)

type producer struct {
	channel *amqp.Channel
}

// 生产者只能给 Exchange 发布信息
func (p *producer) Publish(exchangeName, queueName string, message *brokers.Message) error {
	return nil
}

func newProducer(channel *amqp.Channel) *producer {
	return &producer{
		channel: channel,
	}
}

//func (ch *Channel) Publish(exchange, key string, mandatory, immediate bool, msg Publishing) error

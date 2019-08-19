package rabbitmq

import (
	"sync"

	"github.com/one-hole/simple-rabbitmq/brokers"

	"github.com/streadway/amqp"
)

/************************************************************************************************************************

subscriber 代表一个订阅、通常一个客户端的连接可以产生多个订阅

 ************************************************************************************************************************/

type subscriber struct {
	channel *amqp.Channel
	handler brokers.MessageHandler
}

func newSubscriber(channel *amqp.Channel, handler brokers.MessageHandler) *subscriber {
	return &subscriber{
		channel: channel,
		handler: handler,
	}
}

func (sub *subscriber) run(wg *sync.WaitGroup, queue string) error {

	messages, err := sub.channel.Consume(queue, "", false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case msg := <-messages:
				sub.handler(&brokers.ReceivedMessage{
					ContentType: msg.ContentType,
					Body:        msg.Body,
					MessageID:   msg.MessageId,
					Timestamp:   msg.Timestamp,
				})
				msg.Ack(false)
			}
		}
	}()

	return nil
}

//func (ch *Channel) Publish(exchange, key string, mandatory, immediate bool, msg Publishing) error
//func (ch *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args Table) (<-chan Delivery, error)

//func (c *Connection) Channel() (*Channel, error)

// func (ch *Channel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args Table) error
// func (ch *Channel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error)
// func (ch *Channel) QueueBind(name, key, exchange string, noWait bool, args Table) error

// queue, _ := channel.QueueDeclare(queueName, true, false, false, false, nil)
//	msg, _ := channel.Consume(queue.Name, "", false, false, false, false, nil)

package rabbitmq

import (
	"fmt"
	"github/one-hole/simple-rabbitmq/brokers"
	"sync"

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

	q, err := sub.channel.QueueDeclare(queue, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	messages, err := sub.channel.Consume(q.Name, "", false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case msg := <-messages:
				fmt.Println(string(msg.Body))
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

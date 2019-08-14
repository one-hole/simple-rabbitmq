package rabbitmq

import (
	"errors"
	"github/one-hole/simple-rabbitmq/brokers"

	"github.com/streadway/amqp"
)

type rabbitBroker struct {
	conn                *amqp.Connection
	channelManager      *channelManager      // 信道管理器
	subscriptionManager *subscriptionManager // 订阅管理器
	subscriber          *subscriber          // 这里先放一个订阅
}

func Dial(url string) (brokers.Broker, error) {
	conn, err := amqp.Dial(url)

	if err != nil {
		return nil, err
	}

	channelManager := newChannelManager(conn)
	subscriptionManager := newSubscriptionManager(channelManager)

	return &rabbitBroker{
		conn:                conn,
		channelManager:      channelManager,
		subscriptionManager: subscriptionManager,
	}, nil
}

func (b *rabbitBroker) Close() {
	_ = b.conn.Close()
}

func (b *rabbitBroker) Subscribe(exchange, routingKey string, handler brokers.MessageHandler) error {

	if "" == exchange {
		return errors.New("no exchange error")
	}

	channel, err := b.channelManager.acquireChannel()

	if err != nil {
		panic(err)
	}

	b.subscriber = newSubscriber(channel, handler)

	return b.subscriber.run(nil, routingKey)
}

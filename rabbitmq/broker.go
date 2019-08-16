package rabbitmq

import (
	"github/one-hole/simple-rabbitmq/brokers"

	"github.com/streadway/amqp"
)

type rabbitBroker struct {
	conn                *amqp.Connection
	channelManager      *channelManager      // 信道管理器
	subscriptionManager *subscriptionManager // 订阅管理器
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

// 这里需要区分不同的订阅类型（也就是 ExchangeType）
func (b *rabbitBroker) Subscribe(exchange, routingKey string, handler brokers.MessageHandler) error {

	if "" == exchange {
		return b.subscriptionManager.newSubscription(routingKey, handler)
	}

	if "" == routingKey {
		return b.subscriptionManager.newFanOutSubscription(exchange, handler)
	}

	return nil
}

// 这个地方用了面向接口的 Golang 风格编程
func (b *rabbitBroker) DirectSubscribe(handle *brokers.MessageHandleStruct) error {

	return nil
}

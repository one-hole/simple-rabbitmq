package rabbitmq

import (
	"github.com/one-hole/simple-rabbitmq/brokers"

	"github.com/streadway/amqp"
)

type subscriptionManager struct {
	subscribers    []*subscriber
	channelManager *channelManager
}

// 这里需要是 Topic 的订阅
func (sm *subscriptionManager) newTopicSubscription(exchange string, routingKey string, handler brokers.MessageHandler) error {
	return nil
}

func (sm *subscriptionManager) newDirectSubscription(exchangeName string, handle *brokers.MessageHandleStruct) error {
	channel, _ := sm.channelManager.acquireChannel()
	_ = channel.ExchangeDeclare(exchangeName, amqp.ExchangeDirect, false, false, false, false, nil)
	return nil
}

func (sm *subscriptionManager) newFanOutSubscription(exchangeName string, handler brokers.MessageHandler) error {
	channel, err := sm.channelManager.acquireChannel()

	if err != nil {
		return err
	}

	err = channel.ExchangeDeclare(exchangeName, amqp.ExchangeFanout, true, false, false, false, nil)

	if err != nil {
		return err
	}

	q, err := channel.QueueDeclare("", false, false, true, false, nil)

	if err != nil {
		return err
	}

	err = channel.QueueBind(q.Name, "", exchangeName, false, nil)

	if err != nil {
		return err
	}

	subscriber := newSubscriber(channel, handler)
	subscriber.run(nil, q.Name)
	sm.subscribers = append(sm.subscribers, subscriber)

	return nil
}

func (sm *subscriptionManager) newSubscription(queue string, handler brokers.MessageHandler) error {

	channel, err := sm.channelManager.acquireChannel()
	if err != nil {
		return err
	}

	subscriber := newSubscriber(channel, handler)
	// 理论上这里是不需要创建队列的、但是实际情况就是很多时候 消费者先跑起来了、队列这边的创建也是 find_or_create 的模式
	q, err := channel.QueueDeclare(queue, false, false, false, false, nil)

	if err != nil {
		return err
	}

	subscriber.run(nil, q.Name)

	sm.subscribers = append(sm.subscribers, subscriber)

	return nil
}

func newSubscriptionManager(channelManager *channelManager) *subscriptionManager {
	return &subscriptionManager{
		channelManager: channelManager,
	}
}

//1. Pub/Sub fanout - routing key 失效 ch.QueueBind(queueName, "", exchangeName)
//2. Routing direct
//3. Topic

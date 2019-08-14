package rabbitmq

import "github/one-hole/simple-rabbitmq/brokers"

type subscriptionManager struct {
	subscribers    []*subscriber
	channelManager *channelManager
}

// 这里需要是 Topic 的订阅
func (sm *subscriptionManager) newTopicSubscription(exchange string, routingKey string, handler brokers.MessageHandler) error {
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

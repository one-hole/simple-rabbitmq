package rabbitmq

import "github.com/streadway/amqp"

type channelManager struct {
	conn     *amqp.Connection
	channels []*amqp.Channel
}

func (cm *channelManager) acquireChannel() (*amqp.Channel, error) {
	return cm.conn.Channel()
}

func newChannelManager(conn *amqp.Connection) *channelManager {
	return &channelManager{
		conn:     conn,
		channels: []*amqp.Channel{},
	}
}

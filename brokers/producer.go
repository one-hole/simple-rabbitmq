package brokers

type Producer interface {
	Publish(topic string, message *Message) error
}
package brokers


type Broker interface {
	Close()
	Subscribe(topic, identity string, handler MessageHandler) error
}
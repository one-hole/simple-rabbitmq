package brokers

type Broker interface {
	Close()
	Subscribe(exchange, identity string, handler MessageHandler) error
	DirectSubscribe(handler *MessageHandleStruct) error
}

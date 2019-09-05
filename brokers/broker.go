package brokers

type Broker interface {
	Close()
	DirectSubscribe(handler *MessageHandleStruct) error
	Subscribe(exchange, identity string, handler MessageHandler) error
	//SubscribeV2(exchange string, identity string, exOptions *ExchangeOptions, qOptions *QueueOptions, handler *MessageHandleStruct) error
}

type MQTTBroker interface {
}

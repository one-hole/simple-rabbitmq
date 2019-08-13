package brokers

type MessageHandler func(message *ReceivedMessage) error
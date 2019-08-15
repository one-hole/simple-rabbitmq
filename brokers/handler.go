package brokers

type MessageHandler func(message *ReceivedMessage) error

type MessageHandleStruct interface {
	Handler(message *ReceivedMessage) error
}

// 上面是函数式编程的风格
// 下面式面向接口的编程风格

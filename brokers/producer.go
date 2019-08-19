package brokers

type Producer interface {
	Publish(exchange, queue string, message *Message) error
}

//1. Work Queue | 不需要 Exchange 但是投递的时候需要 Queue 的名字
//2. Publish 不需要 queue name (Fanout Exchange, Exchange Name)

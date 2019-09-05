package brokers

type MessageHandler func(message *ReceivedMessage) error

type MessageHandleStruct interface {
	Handler(message *ReceivedMessage) error
}

/*ExchangeOptions
| Durable | AutoDelete |                             说明                             |
| :-----: | :--------: | :----------------------------------------------------------: |
|  True   |   False    |                  重启之后Exchange会重新定义（默认的）           |
|  True   |    True    | 当 durable 的队列需要绑定在 auto delete 的 exchange 的时候使用  |
|  False  |   False    |          即使没有绑定也不会删除、但是重启不会重定义              |
|  False  |    True    |             当没有绑定的时候会删除、重启不会保留                |
*/
type ExchangeOptions struct {
	Kind       string // RabbitMQ Exchange 的类型 fanout, direct, topic and headers
	Durable    bool   //
	AutoDelete bool   //
	Internal   bool   // 代表不接受 外部的 Publish, exchange 之间投递消息的时候使用
	NoWait     bool   // 这里也就是（定义 Exchange）的时候无需要服务端确认
}

/*QueueOptions
| Durable | AutoDelete |                        说明                         |
| :-----: | :--------: | :-------------------------------------------------:|
|  True   |   False    | 重启之后Queue会重新定义、已经持久化的消息会重新载入 		|
|  True   |    True    |                  基本上不会被使用                    |
|  False  |   False    |     即使没有绑定也不会删除、但是重启不会重定义      	|
|  False  |    True    |        当没有绑定的时候会删除、重启不会保留         	|
*/
type QueueOptions struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	BindingKey bool
}

// 上面是函数式编程的风格
// 下面式面向接口的编程风格

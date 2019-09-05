[main.go](https://github.com/one-hole/simple-rabbitmq/blob/master/main.go) 里面给出了 __Subscribe__ 的使用示例

我在 `package brokers` 里 定义了一个 `Broker` 的接口类型

```go

package brokers

type Broker interface {
	Close()
	Subscribe(exchange, identity string, handler MessageHandler) error
	DirectSubscribe(handler *MessageHandleStruct) error
}


```

同时在 `package rabbitmq` 里 实现了该接口

```go
package rabbitmq

func (b *rabbitBroker) Subscribe(exchange, routingKey string, handler brokers.MessageHandler) error {

	if "" == exchange {
		return b.subscriptionManager.newSubscription(routingKey, handler)
	}

	if "" == routingKey {
		return b.subscriptionManager.newFanOutSubscription(exchange, handler)
	}

	return nil
}
```

目前我们关心的 订阅 只有这两种、后期我们会增加 `direct` 和 `topic` 对于相关概念不理解的可以看[这里](https://github.com/one-hole/simple-rabbitmq/blob/master/brokers/README.md)

最终我们只需要在 handler 里面定义处理逻辑即可、这里也可以是一个 struct （如果处理复杂逻辑）

```go
package brokers

type MessageHandleStruct interface {
	Handler(message *ReceivedMessage) error
}


```

### TODO

这边需要处理的是「补偿机制」
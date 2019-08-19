package main

import (
	"fmt"

	"github.com/one-hole/simple-rabbitmq/brokers"
	"github.com/one-hole/simple-rabbitmq/rabbitmq"
)

var (
	forever <-chan bool
)

func main() {

	broker, err := rabbitmq.Dial("amqp://guest:guest@127.0.0.1:5672/")

	if err != nil {
		panic(err)
	}

	defer broker.Close()

	err = broker.Subscribe("", "hello", func(message *brokers.ReceivedMessage) error {
		fmt.Printf("queue: hello, %s\n", string(message.Body))
		return nil
	})

	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}

	err = broker.Subscribe("logs", "", func(message *brokers.ReceivedMessage) error {
		fmt.Printf("exchange: logs, %s\n", string(message.Body))
		return nil
	})

	<-forever
}

//type MessageHandleStruct interface {
//	Handler(message *ReceivedMessage) error
//}

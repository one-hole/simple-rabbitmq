package main

import (
	"fmt"
	"github/one-hole/simple-rabbitmq/brokers"
	"github/one-hole/simple-rabbitmq/rabbitmq"
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
		fmt.Printf("queue: hello, %s", string(message.Body))
		fmt.Println("-------------------")
		return nil
	})

	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}

	err = broker.Subscribe("logs", "", func(message *brokers.ReceivedMessage) error {
		fmt.Printf("queue: hello, %s", string(message.Body))
		return nil
	})

	<-forever
}

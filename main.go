package main

import (
	"github/one-hole/simple-rabbitmq/brokers"
	"log"
)

var (
	//forever <-chan bool
)

type myHandler struct {

}

// 这里个函数做你想做的任何事情
func (mh *myHandler) Handle(body []byte) error {
	log.Printf("Received a message: %s", body)
	return nil
}

func main() {
	client := brokers.NewRabbitMQConnection()
	client.DirectSubscribe("YourName", &myHandler{})
}

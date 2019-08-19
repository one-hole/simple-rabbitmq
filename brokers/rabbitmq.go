package brokers

import (
	"errors"
	"fmt"
	"log"

	"github.com/one-hole/simple-rabbitmq/config"

	"github.com/streadway/amqp"
)

type RabbitConnection struct {
	conn *amqp.Connection
}

// 你的处理 Struct 必须实现 Handle 接口（你可以在这个接口里面处理你的逻辑）
// 我在 main.go 里面有个例子
type messageHandler interface {
	Handle([]byte) error
}

func NewRabbitMQConnection() *RabbitConnection {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Rabbit.User, config.Rabbit.Password, config.Rabbit.Host, config.Rabbit.Port))

	fmt.Println(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Rabbit.User, config.Rabbit.Password, config.Rabbit.Host, config.Rabbit.Port))

	if err != nil {
		log.Fatal(err)
	}
	return &RabbitConnection{conn: conn}
}

func (client *RabbitConnection) DirectSubscribe(queueName string, handler messageHandler) error {
	if "" == queueName {
		return errors.New("no queue name error")
	}

	channel, _ := client.conn.Channel()
	defer channel.Close()

	queue, _ := channel.QueueDeclare(queueName, true, false, false, false, nil)
	msg, _ := channel.Consume(queue.Name, "", false, false, false, false, nil)

	for d := range msg {
		err := handler.Handle(d.Body)
		if err != nil {
			d.Ack(false) // 这里最好是手动 Ack
		}
	}

	return nil
}

//func (client *RabbitConnection) FanoutSubscribe(bindingRouteKey string) error {
//
//}

//func DirectPublish() {
//
//}

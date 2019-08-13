package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"github/one-hole/simple-rabbitmq/brokers"
	"sync"

	"github.com/streadway/amqp"
)

type subscriber struct {
	channel *amqp.Channel
	handler brokers.MessageHandler
}

func newSubscriber(channel *amqp.Channel, handler brokers.MessageHandler) *subscriber {
	return &subscriber{
		channel: channel,
		handler: handler,
	}
}

func (sub *subscriber) run(ctx context.Context, wg *sync.WaitGroup, queue string) error {

	msgs, err := sub.channel.Consume(queue, "", false, false, false, false, nil)

	if err != nil {
		return errors.New(fmt.Sprintf("%s : cannot attach consume on server", err.Error()))
	}

	wg.Add(1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case msg := <-msgs:
				fmt.Println(msg.Body)
			}
		}
	}()

	//wg.Wait()

	return nil
}

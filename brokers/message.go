package brokers

import "time"

type Message struct {
	ContentType string
	Body        []byte
}

type ReceivedMessage struct {
	ContentType string
	Body        []byte
	MessageID   string
	Timestamp   time.Time
}

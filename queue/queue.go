package queue

import "time"

type Message struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

type Queue interface {
	Add(messageName string, payload []byte) error
	AddMessage(message Message) error
	StartConsuming(size int, pollInterval time.Duration, callback func(Message) error)
}

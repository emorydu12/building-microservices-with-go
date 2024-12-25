package queue

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/adjust/rmq"
)

type RedisQueue struct {
	Queue    rmq.Queue
	name     string
	callback func(Message) error
}

var serialNumberLimit *big.Int

func init() {
	serialNumberLimit = new(big.Int).Lsh(big.NewInt(1), 128)
}

func NewRedisQueue(connectionString string, queueName string) (*RedisQueue, error) {
	connection := rmq.OpenConnection("my service", "tcp", connectionString, 1)
	taskQueue := connection.OpenQueue(queueName)

	return &RedisQueue{Queue: taskQueue, name: queueName}, nil
}

func (r *RedisQueue) Add(messageName string, payload []byte) error {
	queuePayload := Message{Name: messageName, Payload: string(payload)}

	return r.AddMessage(queuePayload)
}

func (r *RedisQueue) AddMessage(message Message) error {
	serialNumber, _ := rand.Int(rand.Reader, serialNumberLimit)
	message.ID = strconv.Itoa(time.Now().Nanosecond()) + serialNumber.String()

	payloadBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	fmt.Println("Add event to queue:", string(payloadBytes))
	if !r.Queue.PublishBytes(payloadBytes) {
		return fmt.Errorf("unable to add message the queue")
	}

	return nil
}

func (r *RedisQueue) StartConsuming(size int, pollInterval time.Duration, callback func(Message) error) {
	r.callback = callback
	r.Queue.StartConsuming(size, pollInterval)
	r.Queue.AddConsumer("RedisQueue_"+r.name, r)
}

func (r *RedisQueue) Consume(delivery rmq.Delivery) {
	fmt.Println("Got event from queue:", delivery.Payload())

	message := Message{}

	if err := json.Unmarshal([]byte(delivery.Payload()), &message); err != nil {
		fmt.Println("Error consming event, unable to deserialise event")
		delivery.Reject()
		return
	}

	if err :=r.callback(message); err != nil {
		delivery.Reject()
		return
	}

	delivery.Ack()
}

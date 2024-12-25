package main

import (
	"log"
	"runtime"
	"time"

	"github.com/emorydu/building-microservices-with-go/queue"
)

func main() {
	log.Println("Starting worker")

	q, err := queue.NewRedisQueue("redis:6379", "test_queue")
	if err != nil {
		log.Fatal(err)
	}

	q.StartConsuming(10, 100*time.Millisecond, func(message queue.Message) error {
		log.Printf("Received message: %v, %v, %v\n", message.ID, message.Name, message.Payload)

		return nil
	})

	runtime.Goexit()
}

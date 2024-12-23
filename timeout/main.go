package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/emorydu/building-microservices-with-go/deadline"
)

func main() {
	switch os.Args[1] {
	case "slow":
		makeNormalRequest()
	case "timeout":
		makeTimeoutRequestForDeadline()
	case "withtimeout":
		makeTimeoutRequestForContext()
	case "withdeadline":
		makeTimeoutRequestForWithCDeadline()
	}
}

func makeTimeoutRequestForWithCDeadline() {
	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go func() {
		slow()
	}()

	for {
		select {
		default:
			time.Sleep(200 * time.Millisecond)
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}
	}
}

func makeTimeoutRequestForContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		slow()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func makeTimeoutRequestForDeadline() {
	dl := deadline.New(1 * time.Second)
	err := dl.Run(func(stopper <-chan struct{}) error {
		slow()
		return nil
	})

	switch err {
	case deadline.ErrTimeout:
		fmt.Println("Timeout")
	default:
		fmt.Println(err)
	}
}

func makeNormalRequest() {
	slow()
}

func slow() {
	for i := 0; i < 100; i++ {
		fmt.Println("Loop:", i)
		time.Sleep(1 * time.Second)
	}
}

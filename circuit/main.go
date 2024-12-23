package main

import (
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/breaker"
)

func main() {
	b := breaker.New(3, 1, 5*time.Second)

	for {
		result := b.Run(func() error {
			time.Sleep(2 * time.Second)
			return fmt.Errorf("Timeout")
		})

		switch result {
		case nil:
			// success!
		case breaker.ErrBreakerOpen:
			fmt.Println("Breaker open")
		default:
			fmt.Println(result)
		}

		time.Sleep(500 * time.Millisecond)
	}
}

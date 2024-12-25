package main

import (
	"io"
	"log"
	"net/http"

	"github.com/emorydu/building-microservices-with-go/queue"
)

type Product struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

func main() {
	q, err := queue.NewRedisQueue("redis:6379", "test_queue")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		err := q.Add("new.product", data)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}

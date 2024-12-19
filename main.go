package main

import (
	"log"
	"net/http"
	"os"

	"github.com/emorydu/building-microservices-with-go/data"
	"github.com/emorydu/building-microservices-with-go/handlers"
)

func main() {
	serverURI := "localhost"
	if os.Getenv("DOCKER_IP") != "" {
		serverURI = os.Getenv("DOCKER_IP")
	}

	store, err := data.NewMongoStore(serverURI)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.Search{DataStore: store}
	log.Println("Server starting on port: 8323")

	log.Fatal(http.ListenAndServe(":8323", &handler))
}

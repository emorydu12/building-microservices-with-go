package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

var natsServer = flag.String("nats", "", "NATS server URI")
var natsClient *nats.Conn

type product struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

func init() {
	flag.Parse()
}

func main() {
	var err error
	natsClient, err = nats.Connect("nats://" + *natsServer)
	if err != nil {
		log.Fatal(err)
	}
	defer natsClient.Close()

	http.DefaultServeMux.HandleFunc("/product", productsHandler)

	log.Println("Starting product write service on port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		insertProduct(w, r)
	}
}

func insertProduct(w http.ResponseWriter, r *http.Request) {
	log.Println(("/insert handler called"))

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	natsClient.Publish("product.inserted", data)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		resp := helloWorldResponse{Message: "Hello World:)"}
		data, err := json.Marshal(resp)
		if err != nil {
			panic("Ooops")
		}

		fmt.Fprint(w, string(data))
	})

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func server() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func main() {
	server()	
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var req helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: fmt.Sprintf("Hello %s:)", req.Name)}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

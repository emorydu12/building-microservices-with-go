package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var req helloWorldRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: fmt.Sprintf("Hello %s:)", req.Name)}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

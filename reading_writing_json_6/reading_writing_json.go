package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	cathandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	req := &helloRequest{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	resp := helloResponse{Message: "Hello " + req.Name + ":)"}

	encoder:=json.NewEncoder(w)
	encoder.Encode(&resp)
}

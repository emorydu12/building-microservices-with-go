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

	handler := newValidationHandler(newHelloWorldHandler())

	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	req := helloRequest{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	h.next.ServeHTTP(rw, r)
}

type helloWorldHandler struct {
}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := helloResponse{Message: "Hello"}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

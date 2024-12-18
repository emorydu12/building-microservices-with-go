package main

import (
	"context"
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

type contextKey string

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return &validationHandler{next: next}
}

func (h *validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &helloWorldRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), contextKey("name"), req.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(w, r)
}

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return &helloWorldHandler{}
}

func (h *helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(contextKey("name")).(string)

	response := helloWorldResponse{Message: "Hello " + name + ":)"}

	_ = json.NewEncoder(w).Encode(response)
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

package handlers

import (
	"net/http"

	"github.com/google/uuid"
)

type correlationHandler struct {
	next http.Handler
}

func (c *correlationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Request-ID") == "" {
		r.Header.Set("X-Request-ID", uuid.New().String())
	}

	c.next.ServeHTTP(w, r)
}

func NewCorrelationHandler(next http.Handler) http.Handler {
	return &correlationHandler{next: next}
}

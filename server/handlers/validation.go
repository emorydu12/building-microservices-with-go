package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/alexcesaro/statsd"
	"github.com/emorydu/building-microservices-with-go/server/entities"
	"github.com/emorydu/building-microservices-with-go/server/httputil"
	"github.com/sirupsen/logrus"
)

type validationHandler struct {
	next   http.Handler
	statsd *statsd.Client
	logger *logrus.Logger
}

func NewValidationHandler(statsd *statsd.Client, logger *logrus.Logger, next http.Handler) http.Handler {
	return &validationHandler{next: next, statsd: statsd, logger: logger}
}

func (h *validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req entities.HelloWorldRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.statsd.Increment(validationFailed)

		message := httputil.SerialzableRequest{Request: r}
		h.logger.WithFields(logrus.Fields{
			"handler": "Validation",
			"status":  http.StatusBadRequest,
			"method":  r.Method,
		}).Info(message.ToJSON())

		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), "name", req.Name)
	r = r.WithContext(c)

	h.statsd.Increment(validationSuccess)

	h.next.ServeHTTP(w, r)
}

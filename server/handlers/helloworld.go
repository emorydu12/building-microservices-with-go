package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/alexcesaro/statsd"
	"github.com/emorydu/building-microservices-with-go/server/entities"
	"github.com/emorydu/building-microservices-with-go/server/httputil"
	"github.com/sirupsen/logrus"
)

type helloWorldHandler struct {
	statsd *statsd.Client
	logger *logrus.Logger
}

func NewHelloWorldHandler(statsd *statsd.Client, logger *logrus.Logger) http.Handler {
	return &helloWorldHandler{statsd: statsd, logger: logger}
}

func (h *helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	timing := h.statsd.NewTiming()

	name := r.Context().Value("name").(string)
	response := entities.HelloWorldResponse{Message: "Hello " + name + ":)"}

	json.NewEncoder(w).Encode(response)

	h.statsd.Increment(helloworldSuccess)

	message := httputil.SerialzableRequest{Request: r}
	h.logger.WithFields(logrus.Fields{
		"handler": "HelloWorld",
		"status":  http.StatusOK,
		"method":  r.Method,
	}).Info(message.ToJSON())

	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

	timing.Send(helloworldTiming)
}

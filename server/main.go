package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/alexcesaro/statsd"
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/emorydu/building-microservices-with-go/server/handlers"
	"github.com/sirupsen/logrus"
)

const port = 8091

func main() {
	statsd, err := createStatsDClient(os.Getenv("STATSD"))
	if err != nil {
		log.Fatal("Unable to create statsD client")
	}

	logger, err := createLogger(os.Getenv("LOGSTASH"))
	if err != nil {
		log.Fatal("Unable to create logstash client")
	}

	setupHandlers(statsd, logger)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func setupHandlers(statsd *statsd.Client, logger *logrus.Logger) {
	validation := handlers.NewValidationHandler(
		statsd,
		logger,
		handlers.NewHelloWorldHandler(statsd, logger),
	)

	bangHandler := handlers.NewPanicHandler(
		statsd,
		logger,
		handlers.NewBangHandler(),
	)

	http.Handle("/helloworld", handlers.NewCorrelationHandler(validation))
	http.Handle("/bang", handlers.NewCorrelationHandler(bangHandler))
}

func createLogger(addr string) (*logrus.Logger, error) {
	retryCount := 0

	log.Println("LOGSTASH:", addr)
	l := logrus.New()
	hostname, _ := os.Hostname()
	var err error

	for ; retryCount < 10; retryCount++ {
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			hook := logrustash.New(
				conn,
				logrustash.DefaultFormatter(
					logrus.Fields{"hostname": hostname},
				),
			)

			l.Hooks.Add(hook)
			return l, err
		}

		log.Println("Unable to connect to logstash, retrying...")
		time.Sleep(1 * time.Second)
	}

	log.Fatal("Unable to connect to logstash")

	return nil, err
}

func createStatsDClient(addr string) (*statsd.Client, error) {
	return statsd.New(statsd.Address(addr))
}

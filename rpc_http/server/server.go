package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/emorydu/building-microservices-with-go/rpc_http/contract"
)

const port = 1234

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name + ":)"

	return nil
}

func StartServer() {
	handler := &HelloWorldHandler{}
	rpc.Register(handler)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %v", port), err)
	}

	log.Printf("Server starting on port: %v", port)

	http.Serve(l, nil)
}

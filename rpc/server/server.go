package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/emorydu/building-microservices-with-go/rpc/contract"
)

const port = 1234

// func main() {
	// log.Printf("Server starting on port: %v\n", port)

	// StartServer()
// }

func StartServer() {
	handler := &HelloWorldHandler{}
	rpc.Register(handler)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %v", port), err)
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name + ":)"

	return nil
}

package main

import (
	"fmt"

	"github.com/emorydu/building-microservices-with-go/rpc_http/client"
	"github.com/emorydu/building-microservices-with-go/rpc_http/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}

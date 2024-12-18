package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/emorydu/building-microservices-with-go/rpc_http_json/contract"
)

func PerformRequest() contract.HelloWorldResponse {
	r, _ := http.Post(
		"http://localhost:1234",
		"application/json",
		bytes.NewBuffer([]byte(`{"id": 1, "method": "HelloWorldHandler", "params": [{"name": "World"}]}`)),
	)
	defer func() {
		_ = r.Body.Close()
	}()

	decoder := json.NewDecoder(r.Body)
	var response contract.HelloWorldResponse
	decoder.Decode(&response)

	return response
}

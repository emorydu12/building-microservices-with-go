package main

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"
)

type Response struct {
	Message string
}

func BenchmarkHelloHandlerVariable(b *testing.B) {
	b.ResetTimer()

	writer := io.Discard
	resp := Response{Message: "Hello World:)"}

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(resp)
		fmt.Fprint(writer, string(data))
	}
}

func BenchmarkHelloHandlerEncoder(b *testing.B) {
	b.ResetTimer()

	w := io.Discard
	resp := Response{Message: "Hello World:)"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(w)
		encoder.Encode(resp)
	}
}

func BenchmarkHelloHandlerEncoderRef(b *testing.B) {
	b.ResetTimer()

	w := io.Discard
	resp := Response{Message: "Hello World:)"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(w)
		encoder.Encode(&resp)
	}
}

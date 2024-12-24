package main

import (
	"fmt"
	"log"
	"net/http"
)

// generate key:
// openssl ecparam -genkey -name secp3841r1 -out server.key

// generate key

// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Hello World:)")
	})

	err := http.ListenAndServeTLS(":8433", "../generate_keys/instance_cert.pem", "../generate_keys/instance_key.pem", nil)

	log.Fatal(err)
}

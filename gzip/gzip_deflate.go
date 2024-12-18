package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var name string

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.Handle("/helloworld",
		NewGzipHandler(http.HandlerFunc(helloWorldHandler)),
	)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello " + name + ":)"}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

type GzipHandler struct {
	next http.Handler
}

func NewGzipHandler(next http.Handler) http.Handler {
	return &GzipHandler{next: next}
}

func (h *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodings := r.Header.Get("Accept-Encoding")

	if strings.Contains(encodings, "gzip") {
		h.serveGzipped(w, r)
	} else if strings.Contains(encodings, "deflate") {
		panic("Deflate not implemented")
	} else {
		h.servePlain(w, r)
	}
}

type GzipResponseWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := w.Header()["Content-Type"]; !ok {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}

	return w.gw.Write(b)
}

func (w GzipResponseWriter) Flush() {
	w.gw.Flush()
	if fw, ok := w.ResponseWriter.(http.Flusher); ok {
		fw.Flush()
	}
}

func (h *GzipHandler) serveGzipped(w http.ResponseWriter, r *http.Request) {
	gzw := gzip.NewWriter(w)
	defer gzw.Close()

	w.Header().Set("Content-Type", "gzip")
	h.next.ServeHTTP(GzipResponseWriter{gzw, w}, r)
}

func (h *GzipHandler) servePlain(w http.ResponseWriter, r *http.Request) {
	h.next.ServeHTTP(w, r)
}

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/VividCortex/ewma"
)

var ma ewma.MovingAverage
var threshold = 1000 * time.Millisecond
var timeout = 1000 * time.Microsecond
var resetting = false
var resetMutex = sync.RWMutex{}

func main() {
	ma = ewma.NewMovingAverage()

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if !isHealthy() {
		respondServiceUnhealthy(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Average request time: %f (ms)\n", ma.Value()/1000000)

	d := time.Since(start)
	ma.Add(float64(d))
}

func respondServiceUnhealthy(w http.ResponseWriter) {
	w.WriteHeader(http.StatusServiceUnavailable)

	resetMutex.RLock()
	defer resetMutex.RUnlock()

	if !resetting {
		go sleepAndResetAverage()
	}
}

func sleepAndResetAverage() {
	resetMutex.Lock()
	resetting = true
	resetMutex.Unlock()

	time.Sleep(timeout)
	ma = ewma.NewMovingAverage()

	resetMutex.Lock()
	resetting = false
	resetMutex.Unlock()
}

func isHealthy() bool {
	return (ma.Value() < float64(threshold))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if !isHealthy() {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	fmt.Fprint(w, "OK")
}

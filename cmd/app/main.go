package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"sync/atomic"
)

const (
	port = ":7092"
)

var views = map[string]string{
	"/":        "/",
	"headers": "/headers",
}

var (
	number = randomFunction()
	counter = int32(0)
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! app %d, %d\n", number, atomic.AddInt32(&counter, 1))
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

func main() {
	v := make([]string, 0, len(views))
	for _, path := range views {
		v = append(v, path)
	}

	fmt.Printf("Starting server on port %s, views: %v\n", port, v)

	http.HandleFunc(views["/"], hello)
	http.HandleFunc(views["headers"], headers)
	http.ListenAndServe(port, nil)
}

func randomFunction() int {
	min := 1
	max := 100
	return rand.Intn(max-min) + min
}
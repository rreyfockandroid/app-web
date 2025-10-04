package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":7092"
)

var views = map[string]string{
	"/":        "/",
	"headers": "/headers",
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
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
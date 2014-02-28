package main

import (
	"flag"
	"fmt"
	"net/http"
)

var addr string

func init() {
	flag.StringVar(&addr, "http", ":4000", "address or port to bind to")
	flag.Parse()
}

func main() {
	r := router()
	http.Handle("/", r)
	fmt.Printf("Listening on %s\n", addr)
	http.ListenAndServe(addr, nil)
} //

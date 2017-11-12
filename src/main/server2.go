package main

import (
	"fmt"
	"net/http"
	"log"
	"sync"
)

var mu sync.Mutex
var count int

func main () {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal( http.ListenAndServe("localhost:8881", nil) )
}

func handler ( w http.ResponseWriter, r *http.Request ) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "Request URL:%q\n", r.URL.Path)
}

func counter ( w http.ResponseWriter, r *http.Request ) {
	mu.Lock()
	fmt.Fprintf(w, "counter: %d\n", count)
	mu.Unlock()
}
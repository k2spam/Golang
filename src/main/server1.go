package main

import (
	"fmt"
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/", httpHandler)
	log.Fatal( http.ListenAndServe("localhost:8888", nil) )
}

func httpHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "url.path=%q\n", r.URL.Path)
}

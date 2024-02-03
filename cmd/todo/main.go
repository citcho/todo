package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", "hello world")
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}

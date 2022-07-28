package main

import (
	"log"
	"net/http"
)

// main function of program
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// starting new web server
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

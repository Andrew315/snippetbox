package main

import (
	"log"
	"net/http"
)

// Home handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// main function of program
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// starting new web server
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

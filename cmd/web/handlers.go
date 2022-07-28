package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home handler
func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

//showSnippet handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Display a specific snippet..."))
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displat a specific snipper with ID %d ...", id)
}

//createSnippet handler function
func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")

		//http.Error intead
		// w.WriteHeader(405)
		// w.Write([]byte(`{"name":"Alex"}`))
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

package main

import (
	"log"
	"net/http"
)

//define route handlers

// home handler
func home(w http.ResponseWriter, r *http.Request) {
	//assertains the path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Let's Go!"))
}

// snippet handler
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is your code snippet!"))
}

// create snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create your code snippet!"))
}

func main() {

	//initialize a server
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//	listen to port
	log.Println("Listening on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

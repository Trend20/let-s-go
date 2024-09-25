## Let's Go

### Path Types

**Fixed path:** has no trailing slash

**Subtree path:** has a trailing slash(/)

To make a check for the paths, just include a check of the URL path inside the handle function.

[//]: # (checks)
It is possible to call `w.WriteHeader()` method just once per response and it should always be called 
before `w.Write()`. This is because if the later is called first, it will send a status code of 200
to the use but if you want to send a status code different from 200, then call the `w.WriteHeader()` before.

`package main
import (
"fmt"
"log"
"net/http"
"strconv"
)
//define route handlers`

`// home handler
func home(w http.ResponseWriter, r *http.Request) {
//assertains the path
if r.URL.Path != "/" {
http.NotFound(w, r)
return
}
w.Write([]byte("Let's Go!"))
}`

`// snippet handler
func showSnippet(w http.ResponseWriter, r *http.Request) {
//to get query params from a URL string
id, err := strconv.Atoi(r.URL.Query().Get("id"))
fmt.Println(id, err)
if err != nil || id < 1 {
http.NotFound(w, r)
return
}
w.Write([]byte("This is your code snippet!"))
}`

`// create snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
        //check for the write HTTP method
        if r.Method != http.MethodPost {
        //set allowed method
        w.Header().Set("Allow", http.MethodPost)
		//other helper functions from w.Header method
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		//add appends the headers
		w.Header().Add("Content-Type", "text/plain")
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")
		//del, deletes the headers
		w.Header().Del("Cache-Control")
		//set status code explicitly
		//w.WriteHeader(405)
		//w.Write([]byte("Method not allowed"))
		//replaces the above method calls and calls them behind the scene
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create your code snippet!"))
}`

`func main() {
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
`
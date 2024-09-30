# Let's Go

<!-- TOC -->
* [Let's Go](#lets-go)
  * [chapter 1: INTRODUCTION](#chapter-1-introduction)
  * [chapter 2: FOUNDATIONS](#chapter-2-foundations)
  * [chapter 3: CONFIGURATION AND ERROR HANDLING](#chapter-3-configuration-and-error-handling)
  * [chapter 4: DATABASE-DRIVEN RESPONSES](#chapter-4-database-driven-responses)
  * [chapter 5: DYNAMIC HTML TEMPLATES](#chapter-5-dynamic-html-templates)
  * [chapter 6: MIDDLEWARE](#chapter-6-middleware)
  * [chapter 7: RESTFUL ROUTING](#chapter-7-restful-routing)
  * [chapter 8: PROCESSING FORMS](#chapter-8-processing-forms)
  * [chapter 9: STATEFUL HTTP](#chapter-9-stateful-http)
  * [chapter 10: SECURITY IMPROVEMENTS](#chapter-10-security-improvements)
  * [chapter 11: USER AUTHENTICATION](#chapter-11-user-authentication)
  * [chapter 12: USING REQUEST CONTEXT](#chapter-12-using-request-context)
  * [chapter 13: TESTING](#chapter-13-testing)
  * [chapter 14: CONCLUSION](#chapter-14-conclusion)
  * [chapter 15: APPENDICES](#chapter-15-appendices)
  * [chapter 16: GUIDED EXERCISES](#chapter-16-guided-exercises)
    * [Path Types](#path-types)
<!-- TOC -->

## chapter 1: INTRODUCTION


## chapter 2: FOUNDATIONS
You can check your Go version after a successful installation by running : `go version`

To create a new project in Go, create a new directory anywhere on your computer and initialize a module
by running: `go mod init <path>`

The module path is just the canonical name or identifier for your project.

To know more about your Go current installation and environment, open the terminal and run: `go env `

To create a web application in Go, we need 3 most essential components:
1. **Handler**: If you're from the MVC background, you can think of them as controllers. They are responsible
for executing your application logic and for writing HTTP response headers and bodies.
2. **Router**: It is also known as _**servemux**_ in Go.
3. **Web server**: This is the last thing that we need and you can always establish a web server and listen to it 
natively without importing any third-party package.

`http.Handle()` and `http.HandleFunc()` allows you to create routes without declaring a servemux.


`http.HandleFunc("/", home) `

`http.HandleFunc("/snippet", showSnippet)`

Behind the scenes, these functions register their routes with something called `DefaultServeMux` which is just a regular
servemux which is initialized by default and stored in the `net/http` package.

## chapter 3: CONFIGURATION AND ERROR HANDLING

## chapter 4: DATABASE-DRIVEN RESPONSES

## chapter 5: DYNAMIC HTML TEMPLATES

## chapter 6: MIDDLEWARE

## chapter 7: RESTFUL ROUTING

## chapter 8: PROCESSING FORMS

## chapter 9: STATEFUL HTTP

## chapter 10: SECURITY IMPROVEMENTS

## chapter 11: USER AUTHENTICATION

## chapter 12: USING REQUEST CONTEXT

## chapter 13: TESTING

## chapter 14: CONCLUSION

## chapter 15: APPENDICES

## chapter 16: GUIDED EXERCISES



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
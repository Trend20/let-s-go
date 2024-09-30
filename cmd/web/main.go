package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	//flag to assign the address of the application
	addr := flag.String("addr", ":4000", "HTTP network address")

	//use environment variable
	//addr := os.Getenv("SNIPPETBOX_ADDR")

	//parse the flag
	flag.Parse()

	//pre existing variables
	//This can be useful if you want to store all your configuration settings in a single struct. As a rough example:
	type Config struct {
		Addr      string
		StaticDir string
	}
	//cfg := new(Config)
	//flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	//flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	//flag.Parse()

	//leveled logging

	//informational message log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//error log
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	//Note that the path given to the http.Dir function is relative to the project
	//directory root.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the mux.Handle() function to register the file server as the handler for
	//all URL paths that start with "/static/". For matching paths, we strip the
	//"/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	//	listen to the post
	infoLog.Printf("Starting server on %S", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

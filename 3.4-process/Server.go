package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

/**
 * Handler for all errors.
 */
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html")
	if status == http.StatusNotFound {
		fmt.Println("Page not found url: " + r.URL.Path)
		fmt.Fprint(w, "<h4>(404) Page Not Found!</h4>")
		return
	}
	fmt.Fprint(w, "<h1>Something went wrong! -- go server</h1>")
}

/**
 * Handles the / request with the html telling what to do.
 */
func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	fmt.Fprintf(w, `<h1>Hello world</h1>`)
}

/**
 * Handles the /pid request, responds with the current process id
 */
func pidHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/pid" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)

	pid := os.Getpid()
	fmt.Fprintf(w, "%v", pid) 
}

// The entry point.
func main() {
	portPtr := flag.String("port", "8000", "Specify the port at which server should listen. e.g. './Server --port=8080'")
	flag.Parse()
	port := *portPtr

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/pid", pidHandler)

	fmt.Println("Server running @ 0.0.0.0:" + port)
	http.ListenAndServe(":"+port, nil)
}

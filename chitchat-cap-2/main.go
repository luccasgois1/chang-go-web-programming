package main

import (
	"net/http"
)

func main() {
	// Create Multiplexer - responsible to read the request url and direct the
	// request to the proper handler of the request
	mux := http.NewServeMux()

	// Serve Static files to the /static route
	// e.g. Requests to http://localhost/static/css/bootstrap.min.css
	// will make the server return the server file located on <application root>/css/bootstrap.min.css
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Root page handler function - All requests for 0.0.0.0:8082/ will be handle by this request
	mux.HandleFunc("/", index)

	// Create the server and start it
	server := &http.Server{
		Addr:    "0.0.0.0:8082",
		Handler: mux,
	}
	server.ListenAndServe()
}

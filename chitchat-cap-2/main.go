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
	// Error page handler function - If a page fails to load for some reason the user should be redirected to this page
	mux.HandleFunc("/err", errHandler)

	// Authentication routes
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)

	// Threads
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// Create the server and start it
	server := &http.Server{
		Addr:    "0.0.0.0:8082",
		Handler: mux,
	}
	server.ListenAndServe()
}

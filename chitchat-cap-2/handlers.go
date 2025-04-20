package main

import (
	"html/template"
	"net/http"

	"github.com/luccasgois1/chang-go-web-programming/chitchat-cap-2/data"
)

// Lets create the handler function for the root route localhost:port/
func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

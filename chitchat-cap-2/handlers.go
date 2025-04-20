package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/luccasgois1/chang-go-web-programming/chitchat-cap-2/data"
)

// Lets create the handler function for the root route localhost:port/
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		generateHTML(w, threads, "layout", "navbar", "index")
	} else {
		redirectToErrorPage(w, r, "Unable to load Threads.")
	}
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	// Get query variables from request to search for the error msg
	vals := r.URL.Query()
	generateHTML(w, vals.Get("msg"), "layout", "navbar", "error")
}

// Generates the base HTML based on the given templates and data to be displayed
func generateHTML(w http.ResponseWriter, data interface{}, filesName ...string) {
	files := []string{}
	for _, fileName := range filesName {
		files = append(files, fmt.Sprintf("templates/%s.html", fileName))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func redirectToErrorPage(w http.ResponseWriter, r *http.Request, errMsg string) {
	redirectUrl := fmt.Sprintf("/err?msg=%s", errMsg)
	http.Redirect(w, r, redirectUrl, http.StatusFound)
}

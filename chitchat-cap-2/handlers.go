package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/luccasgois1/chang-go-web-programming/chitchat-cap-2/data"
)

// Lets create the handler function for the root route localhost:port/
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	typeNavbar := "public.navbar"
	if isUserInSession(w, r) {
		typeNavbar = "private.navbar"
	}
	if err == nil {
		generateHTML(w, threads, "layout", typeNavbar, "index")
	} else {
		redirectToErrorPage(w, r, "Unable to load Threads.")
	}
}

// Handler function to show an error page
func errHandler(w http.ResponseWriter, r *http.Request) {
	// Get query variables from request to search for the error msg
	vals := r.URL.Query()
	typeNavbar := "public.navbar"
	if isUserInSession(w, r) {
		typeNavbar = "private.navbar"
	}
	generateHTML(w, vals.Get("msg"), "layout", typeNavbar, "error")
}

// Authenticates a user given the username and password
func authenticate(w http.ResponseWriter, r *http.Request) {
	// I decided to return this error for the user if
	// any problems in this part of the code will give the user the
	// idea that something is wrong with his credentials. This is
	// a way to protect the app from possible hackers that want to
	// know information about the authentication process.
	errMessageForUser := "Incorrect credentials."
	err := r.ParseForm()
	if err != nil {
		log.Println("failed to parse the login form of the request.", err)
		redirectToErrorPage(w, r, errMessageForUser)
	}

	// Getting user by email
	user := data.UserByEmail(r.PostFormValue("email"))
	if user.IsEmpty() {
		log.Println("failed to get the user from database. user object returned empty.")
		redirectToErrorPage(w, r, errMessageForUser)
	}

	// Check if provided password match the one in the server
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		// Creates a session
		session := data.CreateSession(&user)
		// Check if the session is valid
		if session.IsEmpty() {
			log.Println("failed to create session. the object returned empty")
			redirectToErrorPage(w, r, errMessageForUser)
		}
		// Save the session cookie
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		// Go to the root page but if the session cookie
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// Goes back to the login page
		log.Println("failed to authenticate due to incorrected password.")
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		log.Println("failed to get the cookie", err)
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "signup")
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	errMessageForUser := "unable to create account for this user"
	err := r.ParseForm()
	if err != nil {
		log.Println("failed to parse the signup form.", err)
		redirectToErrorPage(w, r, errMessageForUser)
	}
	user := data.CreateUser(
		r.PostFormValue("name"),
		r.PostFormValue("email"),
		r.PostFormValue("password"))
	if user.IsEmpty() {
		log.Println("failed to create the user on database.")
		redirectToErrorPage(w, r, errMessageForUser)
	}
	http.Redirect(w, r, "/login", http.StatusFound)
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

func isUserInSession(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return false
	}
	session := data.Session{Uuid: cookie.Value}
	return session.IsValid()
}

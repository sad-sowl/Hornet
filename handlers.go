package main

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("GuwJCl9cktJBgTwt8lHuce6gr1UY6u7V")
	store = sessions.NewCookieStore(key)
)

func isValidEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "UserCookie")

	//if person is not authorised  forbid access
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	//if person is authorised then show home page

}

func logup(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "UserCookie")
	session.Values["authenticated"] = false
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/templates/logup.html"))

	if r.Method == "POST" {
		//read data and write in database
		//validate data
		//if everything correct then make session
		session.Values["authenticated"] = true
		session.Save(r, w)
	}

	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "UserCookie")
	tmpl := template.Must(template.ParseFiles("/templates/login.html"))

	if r.Method == "POST" {
		//search in database and validate data

		session.Values["authenticated"] = true
		session.Save(r, w)
	}

	tmpl.Execute(w, nil)
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/sessions"
)

type User struct {
	Name         string
	Number       string
	Passport     string
	Gender       string
	Email        string
	BirthDate    string
	RegAddr      string
	ActualAddr   string
	DeliveryAddr string
}

var (
	db    *sql.DB
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

func isInDatabase(user User) bool {
	query := fmt.Sprintf("SELECT * FROM users WHERE name='%s' OR email='%s' OR number='%s' OR registAddress", user.Name, user.Email, user.RegAddr)

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	if rows.Next() {
		return true
	} else {
		return false
	}
}

func logup(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "UserCookie")
	session.Values["authenticated"] = false
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("logup.html"))

	if r.Method == "POST" {
		//read data and write in the database
		//validate data
		var user User
		user.Name = r.FormValue("name")
		user.Number = r.FormValue("number")
		user.Passport = r.FormValue("passport")
		user.Gender = r.FormValue("gender")
		user.Email = r.FormValue("email")
		user.BirthDate = r.FormValue("data")
		user.RegAddr = r.FormValue("registAddress")
		user.ActualAddr = r.FormValue("actualAddress")
		user.DeliveryAddr = r.FormValue("deliveryAddress")

		if !isValidEmail(user.Email) || isInDatabase(user) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {

			query := fmt.Sprintf("INSERT INTO users(name, number, passport, date, sex, email, registAddress, actualAddress, deliveryAddress) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')",
				user.Name, user.Number, user.Passport, user.BirthDate, user.Gender, user.Email, user.RegAddr, user.ActualAddr, user.DeliveryAddr)

			rows, err := db.Query(query)
			if err != nil {
				panic(err)
			}

			defer rows.Close()

			http.Redirect(w, r, "/thanks", http.StatusTemporaryRedirect)
		}

	}

	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "UserCookie")
	tmpl := template.Must(template.ParseFiles("login.html"))

	if r.Method == "POST" {
		//search in database and validate data

		session.Values["authenticated"] = true
		session.Save(r, w)
	}

	tmpl.Execute(w, nil)
}

package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func main() {

	var err error
	db, err = sql.Open("mysql", "root:Qwerty0106@tcp(127.0.0.1:3306)/hornet")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", logup)
	http.HandleFunc("/login", login)

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

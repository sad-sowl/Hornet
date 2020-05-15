package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", logup)
	http.HandleFunc("/login", login)

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

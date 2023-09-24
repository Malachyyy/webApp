package main

import (
	"fmt"
	"net/http"
	"webApp/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.HandleFunc("/login", handlers.LoginPage)
	http.HandleFunc("/logout", handlers.LogoutPage)

	fmt.Println(fmt.Sprintf("Starting server on: localhost:%s", port))
	_ = http.ListenAndServe(port, nil)
}

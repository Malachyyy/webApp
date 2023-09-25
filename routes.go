package main

import (
	"net/http"
	"webApp/handlers"
)

func routes() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	http.HandleFunc("/login", handlers.LoginPage)
	http.HandleFunc("/signin", handlers.SigninPage)
	http.HandleFunc("/logout", handlers.LogoutPage)
}

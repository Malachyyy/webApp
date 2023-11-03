package handlers

import (
	"io"
	"log"
	"net/http"
	"webApp/functions"
	"webApp/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := render.RenderTemplates(w, "homePage.html")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}

}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := render.RenderTemplates(w, "aboutPage.html")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}

}

func SigninPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := render.RenderTemplates(w, "signInPage.html")

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	err := functions.UsernameValidator(username)
	if err != nil {
		http.Error(w, "Username doesn't meet requirements", http.StatusBadRequest)
	}

	err = functions.PasswordValidator(password)
	if err != nil {
		http.Error(w, "Password doesn't meet requirements", http.StatusBadRequest)
	}

	err = functions.EmailValidator(email)
	if err != nil {
		http.Error(w, "Email doesn't meet requirements", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/login", http.StatusFound)

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}

}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := render.RenderTemplates(w, "loginPage2.html")

	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "1234" {
			io.WriteString(w, "Login successfull")
		} else {
			io.WriteString(w, "Login failed")
		}
	}

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}

}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the logout page\n")
}

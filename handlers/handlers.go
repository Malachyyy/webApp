package handlers

import (
	"io"
	"log"
	"net/http"
	"webApp/pkg"
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

	err := pkg.UsernameValidator(username, password, email)
	if err != nil {
		http.Error(w, "Couldn't validate user", http.StatusBadRequest)
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

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"webApp/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := render.RenderTemplates(w, "homePage.html")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}

}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := render.RenderTemplates(w, "aboutPage.html")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}

}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := render.RenderTemplates(w, "loginPage2.html")

	if r.Method == http.MethodPost {

		login := r.FormValue("login")
		password := r.FormValue("pass")

		if login == "admin" && password == "1234" {
			io.WriteString(w, "Login successfull")
		} else {
			io.WriteString(w, "Login failed")
		}
	}

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}

}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the logout page\n")
}

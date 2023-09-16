package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

const port = ":8080"

func homePage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := renderTemplates(w, "homePage.html")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}

}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := renderTemplates(w, "aboutPage.html")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}

}

func loginPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the login page\n")
}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the logout page\n")
}

func renderTemplates(w http.ResponseWriter, filepath string) *template.Template {
	parsedTemplates, err := template.ParseFiles("./templates/" + filepath)
	if err != nil {
		fmt.Println("Error parsing files:", err)
	}

	return parsedTemplates

}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/logout", logoutPage)

	fmt.Println(fmt.Sprintf("Starting server on: localhost:%s", port))
	_ = http.ListenAndServe(port, nil)
}

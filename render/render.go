package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, filepath string) *template.Template {
	parsedTemplates, err := template.ParseFiles("./templates/" + filepath)
	if err != nil {
		fmt.Println("Error parsing files:", err)
	}

	return parsedTemplates

}

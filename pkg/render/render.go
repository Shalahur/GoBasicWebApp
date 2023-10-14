package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(writer http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(writer, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

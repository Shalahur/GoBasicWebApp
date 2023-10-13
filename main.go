package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func Home(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "home.page.tmpl")
}

func About(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "about.page.tmpl")
}
func renderTemplate(writer http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(writer, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprint("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

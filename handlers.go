package main

import (
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "home.page.tmpl")
}

func About(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "about.page.tmpl")
}

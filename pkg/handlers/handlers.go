package handlers

import (
	"BasicWebApp/pkg/render"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "home.page.tmpl")
}

func About(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "about.page.tmpl")
}

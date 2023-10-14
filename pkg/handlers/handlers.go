package handlers

import (
	"BasicWebApp/pkg/config"
	"BasicWebApp/pkg/render"
	"net/http"
)

// TODO: Repository concept does not understand. need to read more
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "home.page.tmpl")
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "about.page.tmpl")
}

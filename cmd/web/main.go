package main

import (
	"BasicWebApp/pkg/config"
	"BasicWebApp/pkg/handlers"
	"BasicWebApp/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache
	// while developing application make the field value false
	// in production keep the value true
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprint("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

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

var appConfiguration config.AppConfig

func main() {

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	appConfiguration.TemplateCache = templateCache
	// while developing application make the field value false
	// in production keep the value true
	appConfiguration.UseCache = false

	repo := handlers.NewRepo(&appConfiguration)
	handlers.NewHandlers(repo)

	render.NewTemplates(&appConfiguration)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	fmt.Println(fmt.Sprint("Starting application on port %s", portNumber))

	servingPage := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfiguration),
	}

	err = servingPage.ListenAndServe()
	log.Fatal(err)
}

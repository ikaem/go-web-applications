package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ikaem/hello-world-web/pkg/config"
	"github.com/ikaem/hello-world-web/pkg/handlers"
	"github.com/ikaem/hello-world-web/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// Repo is available in the handlers package
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting app on port %s", portNumber)
	// _ = http.ListenAndServe(portNumber, nil) // -> this does not work

	srv := &http.Server{
		Addr:    portNumber, // -> this specifies address - port
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error listening by the server")
	}

}

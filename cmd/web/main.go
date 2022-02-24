package main

import (
	"github.com/OusManDiouf/building-modern-app-with-go/pkg/config"
	"github.com/OusManDiouf/building-modern-app-with-go/pkg/handlers"
	"github.com/OusManDiouf/building-modern-app-with-go/pkg/render"
	"log"
	"net/http"
)

const (
	portNumber = ":8080"
)

func main() {

	var appConfig config.AppConfig

	templateCache, errorTc := render.CreateTemplateCache()
	if errorTc != nil {
		log.Fatal(errorTc)
	}

	appConfig.UseCache = false //devMode
	appConfig.TemplateCache = templateCache

	// Make the appConfig available to the render package
	render.NewTemplates(&appConfig)

	// config the repo
	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	// Handlers
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_ = http.ListenAndServe(portNumber, nil)

}
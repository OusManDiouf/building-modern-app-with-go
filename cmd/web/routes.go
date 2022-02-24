package main

import (
	"github.com/OusManDiouf/building-modern-app-with-go/pkg/config"
	"github.com/OusManDiouf/building-modern-app-with-go/pkg/handlers"
	"github.com/bmizerany/pat"
	"net/http"
)

func routes(appConfig *config.AppConfig) http.Handler {
	mux := pat.New()

	// Setting routes
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}

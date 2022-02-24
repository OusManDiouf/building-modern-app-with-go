package main

import (
	"github.com/OusManDiouf/building-modern-app-with-go/pkg/config"
	"github.com/OusManDiouf/building-modern-app-with-go/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(appConfig *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// Recoverer is a middleware that recovers from panics,
	// logs the panic (and a backtrace),
	// and returns an HTTP 500 (Internal Server Error)
	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

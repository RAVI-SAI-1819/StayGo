package main

import (
	"net/http"
	"github.com/RAVI-SAI-1819/StayGo/pkg/config"
	"github.com/RAVI-SAI-1819/StayGo/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes sets up the application's HTTP routes and middleware
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Recoverer handles panics gracefully
	mux.Use(middleware.Recoverer)

	// Apply CSRF and session middleware
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// Define route handlers
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
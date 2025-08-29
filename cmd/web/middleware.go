package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection using nosurf
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// Configure CSRF cookie settings
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for each request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
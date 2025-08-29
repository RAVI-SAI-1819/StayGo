package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RAVI-SAI-1819/StayGo/pkg/config"
	"github.com/RAVI-SAI-1819/StayGo/pkg/handlers"
	"github.com/RAVI-SAI-1819/StayGo/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// Global app config and session manager
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// Toggle for production mode (affects cookie security)
	app.InProduction = false

	// Initialize session manager with cookie settings
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// Create and assign template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	// Initialize repository and handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// Start HTTP server with configured routes
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Starting application on port number ", portNumber)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
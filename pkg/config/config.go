package config

import (
	"html/template"
	"log"
	"github.com/alexedwards/scs/v2"
)

// AppConfig holds global configuration for the app
type AppConfig struct {
	UseCache      bool                          // Toggle template caching
	TemplateCache map[string]*template.Template // Cached templates
	InfoLog       *log.Logger                   // Optional logger
	InProduction  bool                          // Production mode flag
	Session       *scs.SessionManager           // Session manager
}
package handlers

import (
	"net/http"
	"github.com/RAVI-SAI-1819/StayGo/pkg/config"
	"github.com/RAVI-SAI-1819/StayGo/pkg/models"
	"github.com/RAVI-SAI-1819/StayGo/pkg/render"
)

// Repo is the global repository instance used by handlers
var Repo *Repository

// Repository wraps AppConfig for handler access
type Repository struct {
	App *config.AppConfig
}

// NewRepo initializes a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the global Repo instance
func NewHandlers(r *Repository) {
	Repo = r
}

// Home handler stores remote IP in session and renders home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About handler retrieves remote IP from session and renders about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
package handlers

import (
	"net/http"

	"github.com/karmajigmel/bookings/pkg/config"
	"github.com/karmajigmel/bookings/pkg/models"
	"github.com/karmajigmel/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// SearchIndex is the about page handler
func (m *Repository) SearchIndex(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "searchindex.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// SearchIndex is the about page handler
func (m *Repository) Result(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "result.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

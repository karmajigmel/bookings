package handlers

import (
	"errors"
	"net/http"

	"github.com/karmajigmel/bookings/internal/config"
	"github.com/karmajigmel/bookings/internal/forms"
	"github.com/karmajigmel/bookings/internal/helpers"
	"github.com/karmajigmel/bookings/internal/models"
	"github.com/karmajigmel/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}

// SearchIndex is the about page handler
func (m *Repository) SearchIndex(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "searchindex.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostSearchIndex is the about page handler
func (m *Repository) PostSearchIndex(w http.ResponseWriter, r *http.Request) {
	// name := r.Form.Get("name")
	// dob := r.Form.Get("dob")

	// w.Write([]byte(fmt.Sprintf("Name is %s and date of birth is %s", name, dob)))

	err := r.ParseForm()
	err = errors.New("This is an error message")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		Name: r.Form.Get("name"),
		DOB:  r.Form.Get("dob"),
	}

	form := forms.New(r.PostForm)

	form.Has("dob", r)
	form.Has("name", r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "searchindex.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}

// SearchIndex is the about page handler
func (m *Repository) Result(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "result.page.html", &models.TemplateData{})
}

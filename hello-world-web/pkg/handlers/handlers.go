package handlers

import (
	"net/http"

	"github.com/ikaem/hello-world-web/pkg/config"
	"github.com/ikaem/hello-world-web/pkg/models"
	"github.com/ikaem/hello-world-web/pkg/render"
)

// the repositroy used by the handlers
var Repo *Repository

// the repository type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the about home handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// render.RenderTemplate(w, "home.page.html")
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})

}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

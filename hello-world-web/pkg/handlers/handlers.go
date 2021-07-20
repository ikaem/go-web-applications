package handlers

import (
	"net/http"

	"github.com/ikaem/hello-world-web/pkg/render"
)

// Home is the about home handler
func Home(w http.ResponseWriter, r *http.Request) {

	// render.RenderTemplate(w, "home.page.html")
	render.RenderTemplate(w, "home.page.html")

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.html")

}

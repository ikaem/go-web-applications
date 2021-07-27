package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/ikaem/hello-world-web/pkg/config"
	"github.com/ikaem/hello-world-web/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td

}

// Render function parses and redners an html template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	// _, err := RenderTemplateTest(w)
	// if err != nil {
	// 	fmt.Println("error getting template cache", err)
	// 	return
	// }

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()

	}

	// tc := app.TemplateCache

	// WE dont need to create cache anymore
	// tc, err := CreateTemplateCache()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ok is a boolean
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("No such template")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, AddDefaultData(td))

	_, err := buf.WriteTo(w)

	if err != nil {
		log.Println("Unable to write response from the template buffer")
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	// err = parsedTemplate.Execute(w, nil)
	// log.Print("what is this")

	// if err != nil {
	// 	fmt.Println("error parsing template", err)
	// 	return
	// }
}

// creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		log.Println("thsere was an error")
		return myCache, err
	}

	for _, page := range pages {
		// we use this filepath package to get base path of the page i guess
		// returns last element of the path
		fmt.Println("page is currently:", page)

		name := filepath.Base(page)
		// ts means template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// this is path relative to the root, i guess
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// wee are overwriting the ts . template set object? this is still in the loop
			// parse glob sees to match templates with other tempaltes, based on those directives in the tempaltes
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil
}

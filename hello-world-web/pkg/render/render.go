package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Render function parses and redners an html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	log.Print("what is this")

	if err != nil {
		fmt.Println("error parsing template", err)
	}

}

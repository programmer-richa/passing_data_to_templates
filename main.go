package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// initialise tpl with all the templates available in templates folder with extension gohtml
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	data := map[string]string{"title": "Home", "user": "Richa", "topic": "Programming"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

}

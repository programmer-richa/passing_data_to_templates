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

type Student struct {
	Name          string
	Qualification string
	Grade         string
}

func main() {
	// Passing map to a template
	data := map[string]string{"title": "Home", "user": "Richa", "topic": "Programming"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

	// Passing slice to a template

	languages := []string{"HTML 5", "CSS 3", "Go", "Javascript"}
	err = tpl.ExecuteTemplate(os.Stdout, "languages_slice.gohtml", languages)
	if err != nil {
		log.Fatal(err)
	}

	// Passing map to a template
	student := map[string]string{"Name": "Richa", "Qualification": "MCA", "Grade": "A"}
	err = tpl.ExecuteTemplate(os.Stdout, "student_map.gohtml", student)
	if err != nil {
		log.Fatal(err)
	}

	// Passing struct to a template
	s := Student{
		Name:          "Richa",
		Qualification: "MCA",
		Grade:         "A",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "student_struct.gohtml", s)
	if err != nil {
		log.Fatal(err)
	}

	s1 := Student{
		Name:          "Sahil",
		Qualification: "MCA",
		Grade:         "A",
	}

	students := []Student{s, s1}
	err = tpl.ExecuteTemplate(os.Stdout, "student_slice_of_struct.gohtml", students)
	if err != nil {
		log.Fatal(err)
	}
}

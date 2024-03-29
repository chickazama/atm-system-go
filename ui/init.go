package ui

import (
	"html/template"
	"log"
)

var (
	tmpl *template.Template
)

func init() {
	var err error
	tmpl, err = template.ParseGlob("./templates/*.go.html")
	if err != nil {
		log.Fatal(err.Error())
	}
}

package controllers

import (
	"log"
	"net/http"
	"text/template"
)

var (
	templates *template.Template
)

func init() {
	templates = template.Must(template.ParseFiles("./content/views/index.html"))
}

func renderView(n string, w http.ResponseWriter, model interface{}) {
	err := templates.ExecuteTemplate(w, n+".html", model)
	if err != nil {
		log.Fatal(err)
	}
}

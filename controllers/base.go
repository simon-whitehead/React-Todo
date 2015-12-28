package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/boltdb/bolt"
	"github.com/zenazn/goji/web"
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

func Database(c *web.C) *bolt.DB {
	return c.Env["Database"].(*bolt.DB)
}

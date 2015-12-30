package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/boltdb/bolt"
	"github.com/zenazn/goji/web"
)

var (
	templates map[string]*template.Template
)

func init() {
	templates = make(map[string]*template.Template)
	templates["index"] = template.Must(template.ParseFiles("./content/views/index.html", "./content/views/_base.html"))
}

func renderView(n string, w http.ResponseWriter, model interface{}) {
	err := templates[n].ExecuteTemplate(w, "base", model)
	if err != nil {
		log.Fatal(err)
	}
}

func Database(c *web.C) *bolt.DB {
	return c.Env["Database"].(*bolt.DB)
}

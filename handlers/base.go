package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/boltdb/bolt"
	"github.com/gorilla/sessions"
	"github.com/simon-whitehead/react-todo/middleware"
	"github.com/simon-whitehead/react-todo/services"
	"github.com/zenazn/goji/web"
)

var (
	templates map[string]*template.Template
	funcMap   template.FuncMap
)

func init() {
	templates = make(map[string]*template.Template)
	setupFuncMap()

	templates["index"] = template.Must(template.New("_base").Funcs(funcMap).ParseFiles("./content/views/_base.html", "./content/views/index.html"))
	templates["login"] = template.Must(template.New("_base").Funcs(funcMap).ParseFiles("./content/views/_base.html", "./content/views/login.html"))
}

func renderView(n string, w http.ResponseWriter, model interface{}) {
	err := templates[n].ExecuteTemplate(w, "_base", model)
	if err != nil {
		log.Fatal(err)
	}
}

func setupFuncMap() {
	funcMap = template.FuncMap{
		"any": any,
	}
}

func any(e []interface{}) bool {
	return len(e) > 0
}

func Database(c web.C) *bolt.DB {
	return c.Env["Database"].(*bolt.DB)
}

func Session(c *web.C, r *http.Request) *sessions.Session {
	if s, ok := c.Env["Session"].(*sessions.Session); ok {
		return s
	}

	c.Env["Session"], _ = middleware.NewSession(c, r)
	return c.Env["Session"].(*sessions.Session)
}

func UserService(c web.C) services.UserServicer {
	return c.Env["UserService"].(services.UserServicer)
}

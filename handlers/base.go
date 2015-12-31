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

	templates["index"] = template.Must(template.New("_base").Funcs(funcMap).ParseFiles("./views/_base.html", "./views/index.html"))
	templates["register"] = template.Must(template.New("_base").Funcs(funcMap).ParseFiles("./views/_base.html", "./views/register.html"))
	templates["login"] = template.Must(template.New("_base").Funcs(funcMap).ParseFiles("./views/_base.html", "./views/login.html"))
}

func renderView(c web.C, n string, w http.ResponseWriter, r *http.Request, model interface{}) {
	m := struct {
		Model   interface{}
		Flashes []interface{}
	}{
		Model:   model,
		Flashes: Session(&c, r).Flashes(),
	}
	err := templates[n].ExecuteTemplate(w, "_base", m)
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

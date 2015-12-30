package handlers

import (
	"net/http"

	"github.com/simon-whitehead/react-todo/services"
	"github.com/zenazn/goji/web"
)

func AccountCreatePOST(c web.C, w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	c.Env["UserService"].(services.UserServicer).CreateUser(email, password)

	http.RedirectHandler("/login", http.StatusSeeOther).ServeHTTP(w, r)
}

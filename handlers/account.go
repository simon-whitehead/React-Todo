package handlers

import (
	"net/http"

	"github.com/simon-whitehead/react-todo/domain"
	"github.com/simon-whitehead/react-todo/services"
	"github.com/zenazn/goji/web"
)

func RegisterGET(c web.C, w http.ResponseWriter, r *http.Request) {
	renderView(c, "register", w, r, domain.NewAccountCreateVM())
}

func RegisterPOST(c web.C, w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if _, err := c.Env["UserService"].(services.UserServicer).CreateUser(email, password); err == nil {
		http.RedirectHandler("/login", http.StatusSeeOther).ServeHTTP(w, r)
	} else {
		renderView(c, "register", w, r, domain.NewAccountCreateVM(err))
	}
}

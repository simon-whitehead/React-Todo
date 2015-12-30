package handlers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func LoginIndexGET(c web.C, w http.ResponseWriter, r *http.Request) {
	renderView("login", w, nil)
}

func LoginIndexPOST(c web.C, w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pwd := r.FormValue("password")

	if user, valid := UserService(c).AuthenticateUser(email, pwd); valid {
		session := Session(&c, r)
		session.Values["user"] = user
		err := session.Save(r, w)
		if err != nil {
			fmt.Println("Session save err:", err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		renderView("login", w, struct{ errors string }{errors: "Invalid username/password combination"})
	}
}

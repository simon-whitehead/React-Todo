package handlers

import (
	"fmt"
	"net/http"

	"github.com/simon-whitehead/react-todo/domain"
	"github.com/zenazn/goji/web"
)

func LoginIndexGET(c web.C, w http.ResponseWriter, r *http.Request) {
	renderView(c, "login", w, r, domain.NewLoginVM())
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
		renderView(c, "login", w, r, domain.NewLoginVM("Invalid username/password combination"))
	}
}

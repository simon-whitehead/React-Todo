package handlers

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

func LogoutIndexGET(c web.C, w http.ResponseWriter, r *http.Request) {

	session := Session(&c, r)
	session.Values["user"] = nil
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

package controllers

import (
	"net/http"

	"github.com/simon-whitehead/react-todo/views"
	"github.com/zenazn/goji/web"
)

func HomeIndexGET(c web.C, w http.ResponseWriter, r *http.Request) {
	views.Execute("index", w, nil)
}

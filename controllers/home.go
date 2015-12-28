package controllers

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

func HomeIndexGET(c web.C, w http.ResponseWriter, r *http.Request) {
	renderView("index", w, nil)
}

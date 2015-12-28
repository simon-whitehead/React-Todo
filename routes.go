package main

import (
	"net/http"

	"github.com/simon-whitehead/react-todo/controllers"
	"github.com/zenazn/goji"
)

func init() {
	// Static files (including React code)
	http.Handle("/content/", http.StripPrefix("/content/", http.FileServer(http.Dir("content"))))

	// Main page
	goji.Get("/", controllers.HomeIndexGET)
}

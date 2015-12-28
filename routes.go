package main

import (
	"net/http"

	"github.com/simon-whitehead/react-todo/controllers"
	"github.com/simon-whitehead/react-todo/middleware"
	"github.com/zenazn/goji"
)

func routeSetup() {
	// Static files (including React code)
	http.Handle("/content/", http.StripPrefix("/content/", http.FileServer(http.Dir("content"))))

	goji.Use(middleware.WithDB)

	// Main page
	goji.Get("/", controllers.HomeIndexGET)
}

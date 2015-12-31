package main

import (
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/simon-whitehead/react-todo/handlers"
	"github.com/simon-whitehead/react-todo/middleware"
)

func routeSetup() {
	// Static files (including React code)
	http.Handle("/app/bundle.js", http.StripPrefix("/app/", http.FileServer(http.Dir("app"))))

	// Every route requires the database and a few services
	goji.Use(middleware.Database)
	goji.Use(middleware.UserService)

	// Routes only allowed for authorized users
	var authorizedArea = web.New()

	authorizedArea.Use(middleware.Auth)

	// Main page
	authorizedArea.Get("/", handlers.HomeIndexGET)

	// Anyone can access login and logout
	goji.Get("/register", handlers.RegisterGET)
	goji.Post("/register", handlers.RegisterPOST)

	goji.Get("/login", handlers.LoginIndexGET)
	goji.Post("/login", handlers.LoginIndexPOST)

	goji.Get("/logout", handlers.LogoutIndexGET)

	goji.Handle("/", authorizedArea)
}

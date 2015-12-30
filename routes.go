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
	http.Handle("/content/", http.StripPrefix("/content/", http.FileServer(http.Dir("content"))))

	// Every route requires the database and a few services
	goji.Use(middleware.Database)
	goji.Use(middleware.UserService)

	// Routes only allowed for authorized users
	var authorizedArea = web.New()

	authorizedArea.Use(middleware.Auth)

	// Main page
	authorizedArea.Get("/", handlers.HomeIndexGET)

	// Anyone can access login and logout
	goji.Get("/login", handlers.LoginIndexGET)
	goji.Post("/login", handlers.LoginIndexPOST)
	goji.Get("/logout", handlers.LogoutIndexGET)

	// Anyone can register too
	goji.Post("/account/create", handlers.AccountCreatePOST)

	goji.Handle("/", authorizedArea)
}

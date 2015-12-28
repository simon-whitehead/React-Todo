package main

import (
	"net/http"

	"github.com/simon-whitehead/react-todo/controllers"
	"github.com/simon-whitehead/react-todo/views"
	"github.com/zenazn/goji"
)

func main() {
	views.Initialize()

	http.Handle("/content/", http.StripPrefix("/content/", http.FileServer(http.Dir("content"))))

	goji.Get("/", controllers.HomeIndexGET)

	goji.Serve()
}

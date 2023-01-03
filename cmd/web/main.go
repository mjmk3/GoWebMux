package main

import (
	"GoWebMux/pkg/config"
	"GoWebMux/pkg/handler"
	"GoWebMux/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":5005"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

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

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

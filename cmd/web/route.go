package main

import (
	"GoWebMux/pkg/config"
	"GoWebMux/pkg/handler"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func route(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)

	return mux
}

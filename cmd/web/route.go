package main

import (
	"GoWebMux/pkg/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func route(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	return mux
}

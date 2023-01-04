package main

import (
	"GoWebMux/pkg/config"
	"GoWebMux/pkg/handler"
	"github.com/bmizerany/pat"
	"net/http"
)

func route(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	return mux
}

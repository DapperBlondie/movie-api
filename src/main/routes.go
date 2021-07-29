package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func chiRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/status", config.CheckStatus)

	return mux
}

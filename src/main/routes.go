package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func chiRoutes() http.Handler {
	mux := chi.NewRouter()

	return mux
}
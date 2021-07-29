package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func chiRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/status", config.CheckStatusHandler)
	mux.Get("/movie/:id", config.GetMovieByIDHandler)
	mux.Post("/insert-movie", config.InsertMovieHandler)

	return mux
}

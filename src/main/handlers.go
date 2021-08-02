package main

import (
	"encoding/json"
	"github.com/DapperBlondie/movie-api/src/models"
	"github.com/DapperBlondie/movie-api/src/repo"
	"github.com/go-chi/chi"
	zerolog "github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type MoviePayload struct {
	Name      string   `json:"name"`
	Year      int      `json:"year"`
	Directors []string `json:"directors"`
	Writers   []string `json:"writers"`
	Budget    int64    `json:"budget"`
	Gross     int64    `json:"gross"`
}

type StatusResponse struct {
	Ok      string `json:"ok"`
	Message string `json:"message"`
}

// Config use for holding the application configuration
type Config struct {
	DbRepo *repo.Mongo
}

var config *Config

func NewConfig(dbRepo *repo.Mongo) {
	config = &Config{DbRepo: dbRepo}
	return
}

// CheckStatusHandler use for checking the status of API
func (conf *Config) CheckStatusHandler(w http.ResponseWriter, r *http.Request) {
	resp := &StatusResponse{Ok: "Available", Message: "Everything is OK !"}

	out, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		return
	}
	return
}

// GetMovieByIDHandler use for getting a movie by its own ID
func (conf *Config) GetMovieByIDHandler(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")
	zerolog.Log().Msg(movieID)

	movie := conf.DbRepo.GetMovieByID(movieID)
	w.WriteHeader(http.StatusOK)
	out, err := json.MarshalIndent(movie, "", "\t")
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return
	}

	_, err = w.Write(out)
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return
	}
}

// InsertMovieHandler use for inserting a movie into database
func (conf *Config) InsertMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieP := new(MoviePayload)
	err := json.NewDecoder(r.Body).Decode(movieP)
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return
	}

	movie := &models.Movie{
		ID:        bson.NewObjectIdWithTime(time.Now().UTC()),
		Name:      movieP.Name,
		Year:      movieP.Year,
		Directors: movieP.Directors,
		Writers:   movieP.Writers,
		BOffice: models.BoxOffice{
			Budget: movieP.Budget,
			Gross:  movieP.Gross,
		},
	}

	err = conf.DbRepo.InsertMovie(movie)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
		return
	}
}

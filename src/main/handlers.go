package main

import (
	"github.com/DapperBlondie/movie-api/src/repo"
	"net/http"
)

type StatusReponse struct {
	Ok string `json:"ok"`
}

type Config struct {
	DbRepo *repo.Mongo
}

var config *Config

func NewConfig(dbRepo *repo.Mongo) {
	config = &Config{DbRepo: dbRepo}
}

func (conf *Config) CheckStatus(w http.ResponseWriter, r *http.Request) {

}

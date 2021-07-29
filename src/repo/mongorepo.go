package repo

import (
	"github.com/DapperBlondie/movie-api/src/models"
	zerolog "github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2/bson"
)

func (m *Mongo) GetMovieByID(id string) *models.Movie {
	movie := new(models.Movie)
	err := m.MCollections["movies"].Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&movie)
	if err != nil {
		zerolog.Error().Msg(err.Error() + ", Occurred in GetMovieByID.")
		return nil
	}

	return movie
}

func (m *Mongo) InsertMovie(movie *models.Movie) error {
	err := m.MCollections["movies"].Insert(movie)
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return err
	}
	return nil
}

package repo

import (
	zerolog "github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2"
)

type Mongo struct {
	MSession     *mgo.Session
	MDatabase    *mgo.Database
	MCollections map[string]*mgo.Collection
}

func CreateSession(dsn string) (*Mongo, error) {
	session, err := mgo.Dial(dsn)
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return nil, err
	}

	err = session.Ping()
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return nil, err
	}
	zerolog.Print("Connected to localhost:27017 ...")

	mongo := &Mongo{
		MSession:     session,
		MDatabase:    nil,
		MCollections: make(map[string]*mgo.Collection),
	}

	return mongo, nil
}

func (m *Mongo) AddDataBase(db string) {
	m.MDatabase = m.MSession.DB(db)
}

func (m *Mongo) AddCollection(cName string) {
	m.MCollections[cName] = m.MDatabase.C(cName)
}

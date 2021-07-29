package repo

import (
	zerolog "github.com/rs/zerolog/log"
	"gopkg.in/mgo.v2"
	"time"
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
	defer session.Copy()

	err = session.Ping()
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return nil, err
	}
	zerolog.Print("Connected to localhost:27017 ...")

	session.SetPoolLimit(128)
	session.SetCursorTimeout(time.Minute * 2)
	session.SetSyncTimeout(time.Second * 10)
	session.SetSocketTimeout(time.Minute * 2)

	mongo := &Mongo{MSession: session}

	return mongo, nil
}

func (m *Mongo) AddDataBase(db string) {
	m.MDatabase = m.MSession.DB(db)
}

func (m *Mongo) AddCollection(cName string) {
	m.MCollections[cName] = m.MDatabase.C(cName)
}

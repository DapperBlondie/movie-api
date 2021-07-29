package main

import (
	zerolog "github.com/rs/zerolog/log"
	mgo "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

func main()  {
	session, err := mgo.Dial("localhost")
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return
	}
	defer session.Copy()

	err = session.Ping()
	if err != nil {
		zerolog.Error().Msg(err.Error())
		return
	}
	zerolog.Print("Connected to localhost:27017 ...")

	return
}

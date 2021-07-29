package models

import "gopkg.in/mgo.v2/bson"

type Movie struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Year      int           `json:"year" bson:"year"`
	Directors []string      `json:"directors" bson:"directors"`
	Writers   []string      `json:"writers" bson:"writers"`
	BOffice   BoxOffice     `json:"b_office" bson:"b_office"`
}

type BoxOffice struct {
	Budget int64 `json:"budget" bson:"budget"`
	Gross  int64 `json:"gross" bson:"gross"`
}

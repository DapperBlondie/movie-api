package models

type Movie struct {
	Name string	`json:"name"`
	Year int	`json:"year"`
	Directors []string	`json:"directors"`
	Writers []string	`json:"writers"`
	BOffice BoxOffice	`json:"b_office"`
}

type BoxOffice struct {
	Budget int64	`json:"budget"`
	Gross int64	`json:"gross"`
}

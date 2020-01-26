package models

type Genre struct {
	ID   string `json:_id`
	Name string `json:"name"`
}

var genres []Genre

func init() {
	// populating mock data
	genres = append(genres, Genre{ID: "1", Name: "Action"})
	genres = append(genres, Genre{ID: "2", Name: "Horror"})
}

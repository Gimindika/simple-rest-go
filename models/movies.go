package models

// Model : Movie struct
type Movie struct {
	ID          string `json:"_id"`
	Title       string `json: "title"`
	Description string `json:"description"`
	Genre       Genre  `json:"genre"`
}

// Init movies var as a slive Movie struct - Mock Data
var movies []Movie

func init() {
	// populating mock data
	movies = append(movies, Movie{ID: "1", Title: "Movie One", Description: "Movie One Description", Genre: genres[0]})
	movies = append(movies, Movie{ID: "2", Title: "Movie Two", Description: "Movie Two Description", Genre: genres[1]})
	movies = append(movies, Movie{ID: "3", Title: "Movie Three", Description: "Movie Three Description", Genre: genres[1]})
}

func GetMovies() []Movie {
	return movies
}

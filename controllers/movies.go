package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Gimindika/simple-rest-go/models"
)

type MovieController struct{}

//initialize movie model mock data
var movies = models.GetMovies()

func NewMovieController() *MovieController {
	return &MovieController{}
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (mc MovieController) GetMovies(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	json.NewEncoder(w).Encode(movies)
}

func (mc MovieController) GetMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode(&models.Movie{})
	fmt.Fprint(w, "Movie not found")
}

func (mc MovieController) CreateMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func (mc MovieController) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func (mc MovieController) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

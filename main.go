package main

import (
	"log"
	"net/http"

	"github.com/Gimindika/simple-rest-go/controllers"
	"github.com/gorilla/mux"
)

func main() {
	// initialize controller
	mc := controllers.NewMovieController()

	// initialize router
	r := mux.NewRouter()

	// route handlers
	r.HandleFunc("/api/movies/", mc.GetMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", mc.GetMovie).Methods("GET")
	r.HandleFunc("/api/movies/", mc.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", mc.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", mc.DeleteMovie).Methods("DELETE")

	// start the server
	log.Fatalln(http.ListenAndServe(":5000", r))
}

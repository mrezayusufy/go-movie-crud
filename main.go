package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isdn     string    `json:"isdn"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
	Director *Director `json:"director"`
}
type Director struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Find the movie
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	// send response the movie
	json.NewEncoder(w).Encode(movie)
}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// find the movie and update it
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// find and Delete the movie
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isdn: "32dje", Title: " New Movie", Director: &Director{ID: "1", FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isdn: "3924f", Title: "The key", Director: &Director{ID: "2", FirstName: "Ali", LastName: "Ahmadi"}})
	movies = append(movies, Movie{ID: "3", Isdn: "fdsef9", Title: "The Dragon", Director: &Director{ID: "3", FirstName: "Ahmad", LastName: "Jan"}})
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")
	fmt.Printf("Server is listening to 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

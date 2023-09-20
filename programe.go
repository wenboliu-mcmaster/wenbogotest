package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("conetent-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content_type", "applicaiton/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is to get one movie")
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is to create a movie")
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is to update a movie")

}

func main() {

	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "movie one", Director: &Director{Firstname: "john", Lastname: "snow"}})
	movies = append(movies, Movie{ID: "2", Isbn: "22345", Title: "movie two", Director: &Director{Firstname: "john2", Lastname: "snow2"}})
	r.HandleFunc("/movies", getMovies).Methods("Get")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

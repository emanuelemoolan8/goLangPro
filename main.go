package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, m := range params {
		for k, v := range m {
			fmt.Println(k, "value is", v)
		}
	}

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.Id == params["id"] {
			// json.NewEncoder(w).Encode(movie)
			// return
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	var lastMovieId = movies[len(movies)-1].Id
	intId, err := strconv.Atoi(lastMovieId) // convert string to int
	fmt.Println(intId, err, reflect.TypeOf(intId))
	movie.Id = strconv.Itoa(intId + 1) // convert int to string
	movies = append(movies, movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			fmt.Println(movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
		}
	}
}

//[0,1,2,3,4,5,6,7,8,9]

func main() {
	movies = append(movies, Movie{Id: "1", Isbn: "23456", Title: "Titanic", Director: &Director{Firstname: "James", Lastname: "Camaroon"}})
	movies = append(movies, Movie{Id: "2", Isbn: "65432", Title: "Django", Director: &Director{Firstname: "Jeema", Lastname: "James"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
	// r.run(localhost:8080)
}

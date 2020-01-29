package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"syreclabs.com/go/faker"
)

// Book Model
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Isbn   string  `json:"isbn"`
	Author *Author `json:"author"`
}

// Author Model
type Author struct {
	Name string `json"name"`
	Age  int    `json"age"`
	Sex  string `json"sex"`
}

// The List or Slice bunch of Books
var books []Book

func appendBooks() {
	count := 10
	for index := 0; index < count; index++ {
		books = append(books,
			Book{
				ID:    strconv.Itoa(index),
				Isbn:  strconv.Itoa(123 + index),
				Title: faker.Team().Name(), Author: &Author{
					Name: faker.Name().Name(),
					Age:  rand.Intn(35),
					Sex:  "Female"}})
	}

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	randomInt := strconv.Itoa(rand.Intn(200))
	book.ID = randomInt

	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book

			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"]

			books = append(books, book)

			json.NewEncoder(w).Encode(book)
			return
		}
	}
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}

func main() {

	route := mux.NewRouter()

	appendBooks()

	route.HandleFunc("/api/books", getBooks).Methods("GET")
	route.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	route.HandleFunc("/api/books", createBook).Methods("POST")
	route.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	route.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4545", route))
}

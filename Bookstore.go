package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

//Init books var as a slice Book struct
var books []Book

//Author Struct
type Author struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	//Loop through books to find correct id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100)) //Mock id - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update bok
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] //Mock id - not safe
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func main() {
	// Init Router
	r := mux.NewRouter()

	//Mock data -@todo - implement database
	books = append(books, Book{ID: "1", Isbn: "980123", Title: "Book One", Author: &Author{Firstname: "Donghao", Lastname: "Lyu"}})
	books = append(books, Book{ID: "2", Isbn: "980124", Title: "Book Two", Author: &Author{Firstname: "Wenhui", Lastname: "Xing"}})
	books = append(books, Book{ID: "3", Isbn: "980125", Title: "Book Three", Author: &Author{Firstname: "Fengmei", Lastname: "Wang"}})
	books = append(books, Book{ID: "4", Isbn: "980126", Title: "Book Four", Author: &Author{Firstname: "Shibo", Lastname: "Lyu"}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))
}

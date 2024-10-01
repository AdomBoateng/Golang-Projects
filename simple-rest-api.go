package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Book Structure
type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year string `json:"year"`
}


// Create a slice to hold the books
var books []Book = []Book{
	Book{ID: "1", Title: "Book 1", Author: "Author 1", Year: "2022"},
    Book{ID: "2", Title: "Book 2", Author: "Author 2", Year: "2021"},
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
    _ = json.NewDecoder(r.Body).Decode(&newBook)
    newBook.ID = string(len(books) + 1)
    books = append(books, newBook)
    json.NewEncoder(w).Encode(newBook)
}

func getBooks (w http.ResponseWriter, r *http.Request) {
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
	http.Error(w, "Book not found", http.StatusNotFound)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var updatedBook Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = params["id"]
			books = append(books, updatedBook)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
	http.Error(w, "Book not found", http.StatusNotFound)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"] {
            books = append(books[:index], books[index+1:]...)
            return
        }
    }
    json.NewEncoder(w).Encode(books)
	http.Error(w, "Book deleted", http.StatusNoContent)
}


func main(){
	// Initialize a new router
	r :=mux.NewRouter()

	// Define the routes
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Start the server
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
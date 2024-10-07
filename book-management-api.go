package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

// Add JWT Secret key
var jwtKey = []byte("my_secret_key")

// Book Structure
type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year string `json:"year"`
}

// JWT claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


// Create a slice to hold the books
var books []Book = []Book{
	Book{ID: "1", Title: "Book 1", Author: "Author 1", Year: "2022"},
    Book{ID: "2", Title: "Book 2", Author: "Author 2", Year: "2021"},
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
    err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	
    newBook.ID = string(len(books) + 1)
    books = append(books, newBook)
    json.NewEncoder(w).Encode(newBook)
	respondWithJson(w, http.StatusCreated, newBook)
}

func getBooks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	respondWithJson(w, http.StatusOK, books)
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
	respondWithError(w, http.StatusNotFound, "Book not found")
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
			respondWithJson(w, http.StatusOK, updatedBook)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
	respondWithJson(w, http.StatusNotFound, "Book not found")
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
	respondWithError(w, http.StatusNotFound, "Bad not found")
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
        Password string `json:"password"`
	}

	// Decode 
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if credentials.Username != "admin" || credentials.Password != "password" {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
        return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err!= nil {
        respondWithError(w, http.StatusInternalServerError, "Error creating token")
        return
    }

	respondWithJson(w, http.StatusOK, map[string]string{"token": tokenString})
}

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		tokenString := r.Header.Get("Authorization")

        if tokenString == "" {
            respondWithError(w, http.StatusUnauthorized, "Missing Authorization header")
            return
        }

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err!= nil ||!token.Valid {
            respondWithError(w, http.StatusUnauthorized, "Invalid token")
            return
        }

        next.ServeHTTP(w, r)
	})
}

func main(){
	// Initialize a new router
	r :=mux.NewRouter()

	// Define the routes
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.Handle("/books", jwtMiddleware(http.HandlerFunc(createBook))).Methods("POST")
	r.Handle("/books/{id}", jwtMiddleware(http.HandlerFunc(updateBook))).Methods("PUT")
	r.Handle("/books/{id}", jwtMiddleware(http.HandlerFunc(deleteBook))).Methods("DELETE")

	// Start the server
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
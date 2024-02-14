package bookapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"task6MuxGorm/models"

	"github.com/gorilla/mux"
)

// Create/ Define API Handlers // REST API Standards (TODO: see it)

// POST /books
func (b *bookApi) CreateBook(w http.ResponseWriter, r *http.Request) {

	// TODO: Get JSON data and decode it
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "ERROR READING JSON DATA", http.StatusBadRequest)
	}

	log.Println(book)
	// TODO: call service from here of book
	bk, err := b.Service.CreateBook(&book)

	// SEND RESPONSE
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte(`{"error": "Error creating table"}`))
	}

	// TODO: return json data
	json.NewEncoder(w).Encode(bk)
}

// GET /books
func (b *bookApi) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := b.Service.GetBooks()

	if err != nil {
		w.Write([]byte(`{"error": "Error getting books"}`))
	}

	// TODO: return json data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET /books/{id}
func (b *bookApi) GetOneBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API HANDLED"))
}

// PUT /books/{id}
func (b *bookApi) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// TODO: Get JSON data and decode it
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "ERROR READING JSON DATA", http.StatusBadRequest)
	}

	log.Println(book)

	bk, err := b.Service.UpdateBook(&book)

	// SEND RESPONSE
	if err != nil {
		w.Write([]byte(`{"error": "Error getting books"}`))
	}

	// TODO: return json data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bk)
}

// DELETE /bools/{id}
func (b *bookApi) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	id_int, err := strconv.Atoi(id)

	// GIVE RESPONSE
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte(`{"error": "Error getting id of type uint"}`))
	}
	bk, err := b.Service.DeleteBook(uint(id_int))
	if err != nil {
		w.Write([]byte(`{"error": "Error deleting data of given id"}`))
	}
	json.NewEncoder(w).Encode(bk)
}

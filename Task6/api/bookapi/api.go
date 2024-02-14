package bookapi

import (
	"task6MuxGorm/app"
	"task6MuxGorm/app/bookservice"

	"github.com/gorilla/mux"
)

// define api for books only

// TODO: Should I expose it directly ?? Any other options : Yes i can use interface for it ig in commmon app folder
type bookApi struct {
	App     *app.App
	Service bookservice.Service // From this i can access all services of bookservice
}

// IMP TODO: Do not use pointer here as it will give error like DONT USE POINTER TO INTERFACE

func NewBookApi(a *app.App) *bookApi {

	return &bookApi{
		App:     a,
		Service: bookservice.New(a), // Create New Book Service Here
	}
}

// define routes // starting with "/book/"
func Routes(router *mux.Router, app *app.App) {

	bookApi := NewBookApi(app)

	router.HandleFunc("/", bookApi.CreateBook).Methods("POST")
	router.HandleFunc("/", bookApi.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", bookApi.GetOneBook).Methods("GET")
	router.HandleFunc("/{id}", bookApi.DeleteBook).Methods("DELETE")
	router.HandleFunc("/", bookApi.UpdateBook).Methods("PUT")
}

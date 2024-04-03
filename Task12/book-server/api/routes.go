package api

import "github.com/gorilla/mux"

func (a *api) InitializeRoutes(r *mux.Router) {
	v1Router := r.PathPrefix("/api/v1").Subrouter()

	v1Router.HandleFunc("/book", a.BookApi.HandleGetAllBooks).Methods("GET")

	// Require Authorization
	v1Router.HandleFunc("/book", a.BookApi.HandleCreateNewBook).Methods("POST")
	v1Router.HandleFunc("/book/issue", a.BookApi.HandleIssueOneBook).Methods("POST")
}

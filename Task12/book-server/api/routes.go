package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *api) InitializeRoutes(r *mux.Router) {
	r.Use(mux.MiddlewareFunc(responseJsonMiddleware))
	v1Router := r.PathPrefix("/api/v1").Subrouter()

	v1Router.HandleFunc("/book", a.BookApi.HandleGetAllBooks).Methods("GET")

	// Require Authorization
	v1Router.HandleFunc("/book", a.BookApi.HandleCreateNewBook).Methods("POST")
	v1Router.HandleFunc("/book/issue", a.BookApi.HandleIssueOneBook).Methods("POST")
}

func responseJsonMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

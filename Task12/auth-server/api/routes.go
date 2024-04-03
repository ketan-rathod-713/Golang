package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *api) InitializeRoutes(r *mux.Router) {
	// response type set to json
	r.Use(mux.MiddlewareFunc(responseJsonMiddleware))

	v1Router := r.PathPrefix("/api/v1/").Subrouter()

	v1Router.HandleFunc("/signin", a.AuthApi.HandleSignin).Methods("POST")
	v1Router.HandleFunc("/signup", a.AuthApi.HandleSignUp).Methods("POST")

	v1Router.HandleFunc("/user", a.UserApi.HandleGetAllUsers).Methods("GET")
}

func responseJsonMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

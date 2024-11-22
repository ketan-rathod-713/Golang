package main

import (
	"fmt"
	"log"
	"net/http"
	"wikipediasearch/handlers"
	"wikipediasearch/utils"

	"github.com/gorilla/mux"
)

// define custom unmarshal logic for given data

func main() {
	configs := utils.LoadEnv()

	mx := mux.NewRouter()

	// get search text from user and fetch data from wikepedia api
	mx.HandleFunc("/wikipedia_search", handlers.HandleWikipediaSearch).Methods("POST")
	mx.HandleFunc("/wikipedia_scrap", handlers.HandleWikipediaScrap).Methods("POST")

	// Handle File Server.
	fs := http.FileServer(http.Dir("static"))
	mx.Handle("/", fs)
	// handle all routes
	mx.HandleFunc("/", handlers.CatchAllRouters)

	fmt.Printf("Server Started On Port %v \n", configs.PORT)
	err := http.ListenAndServe(configs.PORT, corsMiddleware(mx))

	if err != nil {
		log.Fatal(err)
	}
}

// Cors middleware: allow request from any origin
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// If it's a preflight request (OPTIONS), respond with success status
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Printf("%v %v \n", r.Method, r.URL.Path)
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

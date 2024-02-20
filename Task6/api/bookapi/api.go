package bookapi

import (
	"task6MuxGorm/app"
	"task6MuxGorm/app/authservice"
	"task6MuxGorm/app/bookservice"

	"github.com/gorilla/mux"
)

type bookApi struct {
	App     *app.App
	Service bookservice.Service // From this i can access all services of bookservice
}

// IMP TODO: Do not use pointer here as it will give error like DONT USE POINTER TO INTERFACE

func newBookApi(a *app.App) *bookApi {

	return &bookApi{
		App:     a,
		Service: bookservice.New(a), // Create New Book Service Here
	}
}

/*
Routes will expose all book api routes with their respective handlers.
TODO: No need to expose whole bookapi itself.
TODO: If any other service required then may be need to expose it too.
*/
func Routes(router *mux.Router, app *app.App) {

	bookApi := newBookApi(app)
	authService := authservice.New(app.Config.JWT_SECRET, app.DB)

	// TODO: only user can access this route

	router.HandleFunc("/", authService.IsAdmin(bookApi.CreateBook)).Methods("POST") // TODO: admin only route
	router.HandleFunc("/", bookApi.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", bookApi.GetOneBookById).Methods("GET")
	router.HandleFunc("/{id}", authService.IsAdmin(bookApi.DeleteBook)).Methods("DELETE") // TODO: admin only route
	router.HandleFunc("/", authService.IsAdmin(bookApi.UpdateBook)).Methods("PUT")        // body should include Id else upsert operation
}

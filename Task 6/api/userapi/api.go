package userapi

import (
	"task6MuxGorm/app"
	"task6MuxGorm/app/userservice"

	"github.com/gorilla/mux"
)

type userApi struct {
	App     *app.App
	Service userservice.Service // From this i can access all services of bookservice
}

// IMP TODO: Do not use pointer here as it will give error like DONT USE POINTER TO INTERFACE

func newUserApi(a *app.App) *userApi {

	return &userApi{
		App:     a,
		Service: userservice.New(a), // Create New Book Service Here
	}
}

/*
Routes will expose all book api routes with their respective handlers.
TODO: No need to expose whole bookapi itself.
TODO: If any other service required then may be need to expose it too.
*/

// /user
func Routes(router *mux.Router, app *app.App) {

	userApi := newUserApi(app)

	router.HandleFunc("/{userId}/book/{bookId}", userApi.IssueBook).Methods("GET") // body should include Id else upsert operation
	router.HandleFunc("/", userApi.CreateUser).Methods("POST")
	router.HandleFunc("/", userApi.GetUsers).Methods("GET")
	router.HandleFunc("/{id}", userApi.GetOneUserById).Methods("GET")
	router.HandleFunc("/{id}", userApi.DeleteUser).Methods("DELETE")
	router.HandleFunc("/", userApi.UpdateUser).Methods("PUT") // body should include Id else upsert operation
}

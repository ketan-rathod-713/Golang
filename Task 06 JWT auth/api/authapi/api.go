package authapi

import (
	"task6MuxGorm/app"
	"task6MuxGorm/app/authservice"

	"github.com/gorilla/mux"
)

type authApi struct {
	App     *app.App
	Service authservice.Service // From this i can access all services of bookservice
}

// IMP TODO: Do not use pointer here as it will give error like DONT USE POINTER TO INTERFACE

func newAuthApi(a *app.App) *authApi {
	return &authApi{
		App:     a,
		Service: authservice.New(a.Config.JWT_SECRET, a.DB), // Create New Book Service Here
	}
}

/*
Routes for auth service
Starts with : /auth/
*/
func Routes(router *mux.Router, app *app.App) {

	authApi := newAuthApi(app)
	// authService := authservice.New(app.Config.JWT_SECRET, app.DB)

	// ! No need of register route as only admin will create it

	router.HandleFunc("/login", authApi.LoginHandler).Methods("POST") // ACCESS: anyone
}

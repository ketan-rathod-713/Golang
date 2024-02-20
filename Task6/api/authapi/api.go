package authapi

import (
	"encoding/json"
	"log"
	"net/http"
	"task6MuxGorm/app"
	"task6MuxGorm/app/authservice"
	"task6MuxGorm/models"
	"task6MuxGorm/utils"
	"time"

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

/* Requres Json Data {userId, password} */
func (a *authApi) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// get credentials // userId and password
	var credentials models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		utils.JSONError(w, utils.ERROR_DECODING_JSON)
		return
	}

	log.Println(credentials)

	// ? validate from database
	var user models.User
	a.App.DB.Where("id = ?", credentials.UserId).Find(&user)

	if user.ID != credentials.UserId {
		utils.JSONError(w, utils.ERROR_FINDING_ID)
		return
	}

	// else check both passwords
	if user.Password != credentials.Password {
		utils.JSONError(w, &models.ApiError{Code: 400, Message: "UserId and password do not match"})
	}

	// Now generate token
	tokenStr, err := a.Service.CreateJwtToken(&user)

	if err != nil {
		log.Println(err)
		utils.JSONError(w, &models.ApiError{Code: 400, Message: "Error Generating Token"})
	}

	// ? Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Path:    "/",
		Expires: time.Now().Add(1 * time.Minute),
	})
}

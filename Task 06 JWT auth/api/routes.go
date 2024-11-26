package api

import (
	"encoding/json"
	"net/http"
	"task6MuxGorm/api/authapi"
	"task6MuxGorm/api/bookapi"
	"task6MuxGorm/api/userapi"
	"task6MuxGorm/app"
	"task6MuxGorm/app/bookservice"

	"github.com/gorilla/mux"
)

type Api struct {
	App         *app.App
	BookService bookservice.Service // API HANDLERS ARE INDEPENDED BUT WE NEED SERVICES HERE SO THAT WE CAN ACCESS IT INSIDE HANDLERS :)
	// TODO: new services will come here
}

func NewApi(app *app.App) (myApi *Api, err error) {
	myApi = &Api{
		App:         app,
		BookService: bookservice.New(app),
	}

	return myApi, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Welcome to task6 gorm crud",
	})
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Invalid URL and METHOD",
	})
}

func (api *Api) InitialiseRoutes(router *mux.Router) {

	// TODO: Define All Handlers Here
	// ALTERNATIVE : Define ALl book related handlers in bookapi and call it from here.

	/* Book Api */
	bookRouter := router.PathPrefix("/book").Subrouter()
	bookapi.Routes(bookRouter, api.App)

	/* User Api */
	userRouter := router.PathPrefix("/user").Subrouter()
	userapi.Routes(userRouter, api.App)

	/* Auth Api */
	authRouter := router.PathPrefix("/auth").Subrouter()
	authapi.Routes(authRouter, api.App)

	// GET / : Home Handler
	router.HandleFunc("/", homeHandler).Methods("GET")

	// Catch All other Url and send invalid response
	router.PathPrefix("/").Handler(http.HandlerFunc(catchAllHandler))
	/* Other Api */
	// other api's will be defined here.
}

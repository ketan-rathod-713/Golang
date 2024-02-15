package api

import (
	"encoding/json"
	"net/http"
	"task6MuxGorm/api/bookapi"
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

func (api *Api) InitialiseRoutes(router *mux.Router) {

	// TODO: Define All Handlers Here
	// ALTERNATIVE : Define ALl book related handlers in bookapi and call it from here.

	// GET / : Home Handler
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "Welcome to task6 gorm crud",
		})

	}).Methods("GET")

	/* Book Api */
	bookRouter := router.PathPrefix("/book").Subrouter()
	bookapi.Routes(bookRouter, api.App)

	/* Other Api */
	// other api's will be defined here.
}

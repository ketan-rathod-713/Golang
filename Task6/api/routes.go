package api

import (
	"task6MuxGorm/api/bookapi"
	"task6MuxGorm/app"
	"task6MuxGorm/app/bookservice"

	"github.com/gorilla/mux"
)

type Api struct {
	App         *app.App
	BookService bookservice.Service // API HANDLERS ARE INDEPENDED BUT WE NEED SERVICES HERE SO THAT WE CAN ACCESS IT INSIDE HANDLERS :)
	// Other Services
}

func NewApi(app *app.App) (myApi *Api, err error) {
	myApi = &Api{
		App:         app,
		BookService: bookservice.New(app),
	}

	return myApi, nil
}

func (api *Api) InitialiseRoutes(router *mux.Router) {
	// Define All Routes
	// Also require DB means define it for particular api instance

	// get particular service information from api instance // as it will be storing all services information // hence i can add routes to that services

	// CREATE NEW BOOK API TO ACCESS ALL HANDLERS ON IT TODO: IMP
	bookApi := bookapi.NewBookApi(api.App)

	router.HandleFunc("/book", bookApi.CreateBook).Methods("POST")
	router.HandleFunc("/book", bookApi.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", bookApi.GetOneBook).Methods("GET")
	router.HandleFunc("/book/{id}", bookApi.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book", bookApi.UpdateBook).Methods("PUT")
}

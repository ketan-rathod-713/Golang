package userapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"task6MuxGorm/models"

	"github.com/gorilla/mux"
)

/*Custom Response Errors */
var ERROR_DECODING_JSON = &models.ApiError{Code: http.StatusBadRequest, Message: "Error Decoding Json Data Provided"}
var ERROR_POST_REQ = &models.ApiError{Code: http.StatusBadRequest, Message: "Error Posting Data for given data"}
var ERROR_GET_REQ = &models.ApiError{Code: http.StatusBadRequest, Message: "Error Fetching Data"}
var ERROR_FINDING_ID = &models.ApiError{Code: http.StatusBadRequest, Message: "Error getting Id of type unsigned int"}

// convert error into json and send it to client
func JSONError(w http.ResponseWriter, err *models.ApiError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}

// POST /user
func (u *userApi) CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Get JSON data and decode it
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		JSONError(w, ERROR_DECODING_JSON)
		return
	}

	log.Println(user)
	// TODO: call service of creating user
	bk, err := u.Service.CreateUser(&user)

	if err != nil {
		JSONError(w, ERROR_POST_REQ)
		return
	}

	// TODO: return json data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bk)
}

// GET /user
func (u *userApi) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.Service.GetUsers()

	if err != nil {
		JSONError(w, ERROR_GET_REQ)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /user/{id}
func (u *userApi) GetOneUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		JSONError(w, ERROR_FINDING_ID)
		return
	}

	user, err := u.Service.GetOneUserById(uint64(id))

	if err != nil {
		JSONError(w, ERROR_GET_REQ)
		return // else bad me null bhi aayega
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// PUT /user/{id}
func (u *userApi) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		JSONError(w, ERROR_DECODING_JSON)
		return
	}

	log.Println(user)

	bk, err := u.Service.UpdateUser(&user)

	if err != nil {
		JSONError(w, &models.ApiError{Code: 500, Message: "Error Updating Data For given id"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bk)
}

// DELETE /user/{id}
func (u *userApi) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	id_int, err := strconv.Atoi(id)

	if err != nil {
		JSONError(w, ERROR_DECODING_JSON)
		return
	}
	bk, err := u.Service.DeleteUser(uint64(id_int))
	if err != nil {
		JSONError(w, &models.ApiError{Code: 500, Message: "Error Deleting Data For given id"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bk)
}

// ! Currently Any User Can Issue Book Ha ha
// POST /book/
func (u *userApi) IssueBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	userId := mux.Vars(r)["userId"]

	bookId_int, err := strconv.Atoi(bookId)
	userId_int, err := strconv.Atoi(userId)

	if err != nil {
		JSONError(w, ERROR_DECODING_JSON)
		return
	}

	bk, err := u.Service.IssueBook(uint64(userId_int), uint64(bookId_int))
	if err != nil {
		JSONError(w, &models.ApiError{Code: 500, Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bk)
}

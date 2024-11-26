package bookapi

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

// @Summary Creates a book
// @BasePath /book/
// @Description To create a book
// @ID book
// @Produce json
// @Success 200 {string} string "ok"
// @Router /book/ [post]
// @Tags book
// @Param data body models.CreateBook true "Book Data"
func (b *bookApi) CreateBook(w http.ResponseWriter, r *http.Request) {
	// TODO: Get JSON data and decode it
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		JSONError(w, ERROR_DECODING_JSON)
		return
	}

	log.Println(book)
	// TODO: call service of creating book
	bk, err := b.Service.CreateBook(&book)

	if err != nil {
		JSONError(w, ERROR_POST_REQ)
		return
	}

	// TODO: return json data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bk)
}

// @Summary Get All Books
// @Description Get All Books
// @ID get-all-books
// @Tags book
// @Produce json
// @Router /book/ [get]
func (b *bookApi) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := b.Service.GetBooks()

	if err != nil {
		JSONError(w, ERROR_GET_REQ)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// @Summary Get a book by id
// @Description Get a book by its id
// @ID get-book-by-id
// @Tags book
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /book/{id} [get]
func (b *bookApi) GetOneBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		JSONError(w, ERROR_FINDING_ID)
		return
	}

	book, err := b.Service.GetOneBookById(uint64(id))

	if err != nil {
		JSONError(w, ERROR_GET_REQ)
		return // else bad me null bhi aayega
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// @Summary Update Book By Id
// @Description update a book by id
// @ID update-book-by-id
// @Tags book
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /book/{id} [get]
func (b *bookApi) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		JSONError(w, ERROR_DECODING_JSON)
		return
	}

	log.Println(book)

	bk, err := b.Service.UpdateBook(&book)

	if err != nil {
		JSONError(w, &models.ApiError{Code: 500, Message: "Error Updating Data For given id"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bk)
}

// @Summary Delete a book by id
// @Description Delete a book by its id
// @ID delete-book-by-id
// @Tags book
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /book/{id} [delete]
func (b *bookApi) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	id_int, err := strconv.Atoi(id)

	if err != nil {
		JSONError(w, ERROR_DECODING_JSON)
		return
	}
	bk, err := b.Service.DeleteBook(uint64(id_int))
	if err != nil {
		JSONError(w, &models.ApiError{Code: 500, Message: "Error Deleting Data For given id"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bk)
}

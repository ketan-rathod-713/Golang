package books

import (
	"books/grpcclient"
	"books/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type errorResponse struct {
	Errors []Error `json:"errors"`
	Type   string  `json:"type"`
}

type Error struct { // can be validation error, internal error etc.
	Field string `json:"field"`
	Error string `json:"error"`
}

func (a *Api) HandleCreateNewBook(w http.ResponseWriter, r *http.Request) {
	// create new book
	w.Header().Set("Content-Type", "application/json")

	type bookCreateRequest struct {
		Title    string  `json:"title" validate:"required" bson:"title,omitempty"`
		Author   string  `json:"author" validate:"required" bson:"author,omitempty"`
		Qty      int     `json:"qty" validate:"required" bson:"qty,omitempty"`
		Price    float64 `json:"price" validate:"required" bson:"price,omitempty"`
		Category string  `json:"category" validate:"required" bson:"category,omitempty"` // Action, Adventure etc.
		Status   string  `json:"status" validate:"required" bson:"status,omitempty"`

		// no need of below fields in request
		// Created   string  `json:"created" validate:"required" bson:"created,omitempty"`
		// Updated   string  `json:"updated" validate:"required" bson:"updated,omitempty"`
		// IsDeleted bool    `json:"isDeleted" bson:"isDeleted"`
	}

	var bookCreateReq bookCreateRequest
	json.NewDecoder(r.Body).Decode(&bookCreateReq)

	// validate all the fields
	err := validate.Struct(&bookCreateReq)

	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		var errors []Error
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Error())
			fmt.Println(err.Field())

			newError := Error{Field: err.Field(), Error: err.Error()}
			errors = append(errors, newError)
		}

		response := errorResponse{Errors: errors, Type: "validation"}
		json.NewEncoder(w).Encode(response)
		return
	} else {

		// get query param token
		token := r.URL.Query().Get("token")

		// Authorise the given user
		authUse := grpcclient.NewAuthUse(a.App.AuthGrpcClient)

		// need jwt token here
		response, err := authUse.AuthoriseUser(token)
		if err != nil {
			fmt.Println("Error Occured In GRPC server", err)
			response := errorResponse{Errors: []Error{Error{Field: "auth", Error: err.Error()}}, Type: "authentication"}
			json.NewEncoder(w).Encode(response)
			return
		}
		fmt.Println("object id", response.ObjectId)

		if err != nil {
			response := errorResponse{Errors: []Error{Error{Field: "auth", Error: err.Error()}}, Type: "authentication"}
			json.NewEncoder(w).Encode(response)
			return
		} else {

			// TODO : We should not take Role data from jwt token as it may be old value in case if it is updated in database
			// TODO : auth service should validate this profile from this point of view too.
			// Now at the end create a document if role is Admin
			if response.Role != "admin" {
				response := errorResponse{Errors: []Error{Error{Field: "role", Error: "you don't have required priviledge to create new book"}}, Type: "authentication"}
				json.NewEncoder(w).Encode(response)
				return
			} else {

				createdBy, err := primitive.ObjectIDFromHex(response.ObjectId)

				if err != nil {
					response := errorResponse{Errors: []Error{Error{Field: "auth", Error: "error fetching user object id from jwt claims" + err.Error()}}, Type: "authentication"}
					json.NewEncoder(w).Encode(response)
					return
				}

				var book models.Book = models.Book{
					Author:    bookCreateReq.Author,
					Title:     bookCreateReq.Title,
					Qty:       bookCreateReq.Qty,
					Category:  bookCreateReq.Category,
					Price:     bookCreateReq.Price,
					Created:   time.Now().String(),
					Updated:   time.Now().String(),
					IsDeleted: false,
					Image:     "no image for now",
					CreatedBy: &createdBy,
				}

				result, err := a.App.DB.Collection("books").InsertOne(context.TODO(), book)

				if err != nil {
					response := errorResponse{Errors: []Error{Error{Field: "book", Error: "Error creating new book"}}, Type: "internal"}
					json.NewEncoder(w).Encode(response)
					return
				}

				switch id := result.InsertedID.(type) {
				case *primitive.ObjectID:
					book.ObjectId = id
				case primitive.ObjectID:
					book.ObjectId = &id
				}

				json.NewEncoder(w).Encode(&book)
			}

		}

	}

}

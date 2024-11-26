package books

import (
	"books/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	v1 "proto/auth-server/v1"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookIssueRequest struct {
	BookId string `json:"bookId" validate:"required,len=24"`
	UserId string `json:"userId" validate:"required,len=24"`
}

type bookIssueResponseJson struct {
	BookIssueRequest BookIssueRequest `json:"bookIssueRequest"`
	Issued           bool             `json:"issued"`
}

// authorize and issue one book on behalf of it.
func (a *Api) HandleIssueOneBook(w http.ResponseWriter, r *http.Request) {
	// For issueing book need to authorise user
	var requestData BookIssueRequest
	json.NewDecoder(r.Body).Decode(&requestData)

	// validate both
	err := validate.Struct(requestData)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		var errors []models.Error
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Error())
			fmt.Println(err.Field())

			newError := models.Error{Field: err.Field(), Error: err.Error()}
			errors = append(errors, newError)
		}

		response := models.ErrorResponse{Errors: errors, Type: "validation"}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		// TODO lock table
		// Now find book from document and check if qty is greater then 0

		// authorise user

		token := r.URL.Query().Get("token")

		authoriseResponse, err := a.App.AuthGrpcClient.AuthoriseUser(context.Background(), &v1.AuthoriseRequest{
			JwtToken: token,
		})

		if err != nil {
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "error", Error: err.Error()}}, Type: "validation"}
			json.NewEncoder(w).Encode(response)
			return
		}

		log.Println(authoriseResponse)

		var book models.Book

		bookId, err := primitive.ObjectIDFromHex(requestData.BookId)
		if err != nil {
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "bookId", Error: err.Error()}}, Type: "validation"}
			json.NewEncoder(w).Encode(response)
			return
		}
		userId, err := primitive.ObjectIDFromHex(requestData.UserId)
		if err != nil {
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "userId", Error: err.Error()}}, Type: "validation"}
			json.NewEncoder(w).Encode(response)
			return
		}

		fmt.Println(userId)

		singleResult := a.App.DB.Collection("books").FindOne(context.TODO(), bson.M{"_id": bookId})
		err = singleResult.Decode(&book)
		if err != nil {
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "error", Error: "Error decoding book from database"}}, Type: "internal"}
			json.NewEncoder(w).Encode(response)
			return
		}

		if book.Qty <= 0 {
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "error", Error: "Not sufficient books to issue"}}, Type: "books"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// update book qty one less
		bookUpdateResult, err := a.App.DB.Collection("books").UpdateByID(context.TODO(), bookId, bson.M{"$set": bson.M{"qty": book.Qty - 1}})

		if err != nil {
			fmt.Println("book update error", err)

			return
		}

		if bookUpdateResult.ModifiedCount != 1 {
			fmt.Println("update book count error")
		}

		bookIssueResponse, err := a.App.AuthGrpcClient.BookIssue(context.TODO(), &v1.BookIssueRequest{
			UserJwtToken: token,
			BookId:       requestData.BookId,
		})

		if err != nil {
			fmt.Println("grpc error")
			return
		}

		if bookIssueResponse.Issued != true {
			fmt.Println("book not issued, cancel the transaction or session")
			return
		}

		// book issued ha ha hence give response to user and commit this session

		response := bookIssueResponseJson{
			Issued:           true,
			BookIssueRequest: requestData,
		}

		json.NewEncoder(w).Encode(response)
	}
}

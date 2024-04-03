package authapi

import (
	"auth/app/jwtauth"
	"auth/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// upon signin find data from database and match it
// then generate jwt token for user and send it as a json.

// for signin only get emailId and Password

func (a *Api) HandleSignin(w http.ResponseWriter, r *http.Request) {
	var signInRequest models.SignInRequest
	json.NewDecoder(r.Body).Decode(&signInRequest)

	// set response type to json
	w.Header().Set("Content-Type", "application/json")

	// Validate Request
	err := validate.Struct(&signInRequest)

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
			newError := models.Error{Field: err.Field(), Error: err.Error()}
			errors = append(errors, newError)
		}

		response := models.ErrorResponse{Errors: errors, Type: "validation"}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		// check for one document with given email Id and Password
		var user models.User
		result := a.DB.Collection("users").FindOne(context.TODO(), bson.M{"email": signInRequest.Email})

		// decode user data
		err = result.Decode(&user)

		if err != nil {
			fmt.Println(err)

			// if no documents found means user does not exist
			if errors.Is(err, mongo.ErrNoDocuments) {
				response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "signin", Error: "User does not exist for given email Id"}}, Type: "internal"}
				json.NewEncoder(w).Encode(response)
				return
			}

			// return err response
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "signin", Error: err.Error()}}, Type: "internal"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// check if password matching or not
		if user.Password != signInRequest.Password {
			fmt.Println("Password Not Matching")
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "signin", Error: "Password Not Matching"}}, Type: "internal"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// generate auth token for the user
		userClaims := jwtauth.UserClaims{ObjectId: *user.ObjectId, Name: user.Name, Email: user.Email, Role: user.Role}
		JwtToken, err := a.JwtService.GenerateJwtToken(&userClaims, a.Config)
		if err != nil {
			fmt.Println(err)
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "signin", Error: err.Error()}}, Type: "internal"}
			json.NewEncoder(w).Encode(response)
			return
		}

		user.JwtToken = JwtToken

		json.NewEncoder(w).Encode(user)
		return
	}
}

// How to send null values if that field is not exists in struct field // todo

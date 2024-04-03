package authapi

import (
	"auth/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Validate Request
var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

// Handle signup will get user info in json formate and it will store it inside database
// encrypt password using hashing // Todo

func (a *Api) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	var req models.SignupRequest
	json.NewDecoder(r.Body).Decode(&req)

	fmt.Println(req)

	err := validate.Struct(req)

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
		var userGot models.User
		// check if given email id already registered or not
		singleResult := a.DB.Collection("users").FindOne(context.TODO(), bson.M{"email": req.Email})
		err = singleResult.Decode(&userGot)

		// if document present with given id then return error
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println(err)
			// return err response
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "signup", Error: err.Error()}}, Type: "internal"}
			json.NewEncoder(w).Encode(response)
			return
		}

		if userGot.Email == req.Email {
			fmt.Println("email id already registered")
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "email", Error: "Email Id Already Registered"}}, Type: "internal"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// save this data to mongodb
		// req.Role = "user"
		result, err := a.DB.Collection("users").InsertOne(context.TODO(), req)

		if err != nil {
			fmt.Println("Error inserting data to mongodb")
			response := models.ErrorResponse{Errors: []models.Error{models.Error{Field: "signup", Error: "Error Inserting data in mongodb database"}}, Type: "internal"}
			json.NewEncoder(w).Encode(response)
			return
		}

		var user models.User

		switch id := result.InsertedID.(type) {
		case primitive.ObjectID:
			user.ObjectId = &id
		}

		// put all data in user and send it

		user.Name = req.Name
		user.Email = req.Email
		user.Phone = req.Phone
		user.Address = req.Address
		user.City = req.City
		user.State = req.State
		user.Country = req.Country
		user.Zip = req.Zip
		user.Standard = req.Standard

		// send user details
		json.NewEncoder(w).Encode(user)
		return
	}
}

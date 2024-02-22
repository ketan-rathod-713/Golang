package userapi

import (
	"fibermongoapp/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

func (u *userApi) CreateUser(ctx *fiber.Ctx) error {
	// get user json data
	var user *models.User
	err := ctx.BodyParser(&user)
	if err != nil { // TODO how to handler errors of handlers in fiber
		return ctx.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "Error Parsing Data", Data: nil})
	}

	//TODO use the validator library to validate required fields
	err = u.Validate.Struct(user)
	if err != nil {
		log.Info("Error validating user")
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Info(err)
			return err
		}

		// To print all validation errors // TODO what is happening here
		for _, err := range err.(validator.ValidationErrors) {
			log.Info(err.Namespace())
			log.Info(err.Field())
			log.Info(err.Tag())
		}

		return ctx.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "An Error Occured", Data: &fiber.Map{"errors": err.Error()}})
	}

	// pass it to service
	user, err = u.Service.CreateUser(user)

	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: &fiber.Map{"error": err.Error()}})
		return err
	}
	// generate response and send

	return ctx.Status(http.StatusOK).JSON(UserResponse{Status: http.StatusOK, Message: "User Created", Data: &fiber.Map{"user": user}})
}

func (u *userApi) GetUsers(ctx *fiber.Ctx) error {

	users, err := u.Service.GetUsers()

	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: nil})
	}

	return ctx.Status(http.StatusOK).JSON(UserResponse{Status: http.StatusOK, Message: "Retrived All Users", Data: &fiber.Map{"users": users}})
}

func (u *userApi) GetOneUserById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	userObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "Error Getting id Param", Data: &fiber.Map{"error": err.Error()}})
	}

	user, err := u.Service.GetOneUserById(userObjectId)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "No documents found", Data: nil})
	}

	return ctx.Status(http.StatusOK).JSON(UserResponse{Status: http.StatusOK, Message: "Retrived One User", Data: &fiber.Map{"users": user}})
}

func (u *userApi) UpdateUser(ctx *fiber.Ctx) error {
	// get data using bodyparser
	var user *models.User
	err := ctx.BodyParser(&user)
	if err != nil { // TODO how to handler errors of handlers in fiber
		return ctx.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "Error Parsing Data", Data: &fiber.Map{"error": err.Error()}})
	}

	// get Id from parameters
	id := ctx.Params("id")

	user.ID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "Error Getting id Param", Data: &fiber.Map{"error": err.Error()}})
	}

	// TODO Validate User Data
	// Here i want only id and any field should be updated how to do it ?

	// call update user service
	user, err = u.Service.UpdateUser(user)
	if err != nil {
		return ctx.Status(http.StatusOK).JSON(UserResponse{Status: http.StatusOK, Message: "Error Updating User", Data: &fiber.Map{"error": err.Error()}})
	}
	// give response

	return ctx.Status(http.StatusOK).JSON(UserResponse{Status: http.StatusOK, Message: "User Updated", Data: &fiber.Map{"updated_user": user}})
}

func (u *userApi) DeleteUser(ctx *fiber.Ctx) error {
	// get id
	id := ctx.Params("id")
	// call service
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(UserResponse{Status: http.StatusBadRequest, Message: "Error Getting id Param", Data: &fiber.Map{"error": err.Error()}})
	}

	user, err := u.Service.DeleteUser(objectId)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "Error Deleting User", Data: &fiber.Map{"error": err.Error()}})
	}

	// return response
	return ctx.Status(http.StatusOK).JSON(UserResponse{Status: http.StatusOK, Message: "User Deleted", Data: &fiber.Map{"deleted_user": user}})
}

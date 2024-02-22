package userapi

import (
	"fibermongoapp/app/userservice"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type userApi struct {
	DB       *mongo.Client
	Service  userservice.Service
	Validate *validator.Validate
}

func New(db *mongo.Client) *userApi {
	return &userApi{DB: db, Service: userservice.New(db), Validate: validator.New(validator.WithRequiredStructEnabled())}
}

func (u *userApi) Routes(parentRouter fiber.Router) {
	parentRouter.Get("/", u.GetUsers)  // TODO add pagination here skip and page will be given
	parentRouter.Get("/:id", u.GetOneUserById)
	parentRouter.Post("/", u.CreateUser)
	parentRouter.Delete("/:id", u.DeleteUser)
	parentRouter.Put("/:id", u.UpdateUser)

	// TODO all other routes don't accept
}

package userapi

import (
	"fibermongoapp/app/userservice"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)
/* wrapper around all user related api's */
type userApi struct {
	DB       *mongo.Client
	Service  userservice.Service
	Validate *validator.Validate
}

func New(db *mongo.Client) *userApi {
	return &userApi{DB: db, Service: userservice.New(db), Validate: validator.New(validator.WithRequiredStructEnabled())}
}

func (u *userApi) Routes(parentRouter fiber.Router) {
	parentRouter.Get("/", u.GetUsers) // TODO: add complex filters such as $gte or $lte etc.
	parentRouter.Get("/:id", u.GetOneUserById)
	parentRouter.Post("/", u.CreateUser)
	parentRouter.Delete("/:id", u.DeleteUser)
	parentRouter.Put("/:id", u.UpdateUser)

}

package user

import (
	"context"
	"graphql_search/graph/model"
	"graphql_search/models"
	"graphql_search/service/auth"
	"graphql_search/service/mail"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type api struct {
	Database       *mongo.Database
	DB_Collections *models.DB_COLLECTIONS
	Validator      *validator.Validate
	AuthService    auth.Service
	MailService    mail.Service
}

func New(database *mongo.Database, dbCollections *models.DB_COLLECTIONS, authService auth.Service, mailSerive mail.Service) Api {
	return &api{
		Database:       database,
		DB_Collections: dbCollections,
		Validator:      validator.New(validator.WithRequiredStructEnabled()),
		AuthService:    authService,
		MailService:    mailSerive,
	}
}

type Api interface {
	RegisterUser(ctx context.Context, userInput model.RegisterUser) (*models.User, error)
	SignInUserByEmail(ctx context.Context, user model.SignInUserByEmail) (*models.User, error)
	GetUser(ctx context.Context, authToken string) (*models.User, error)
	VerifyUserEmail(ctx context.Context, authToken string) (*models.User, error)
}

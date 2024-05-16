package user

import (
	"context"
	"errors"
	"fmt"
	"graphql_search/graph/model"
	"graphql_search/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *api) RegisterUser(ctx context.Context, userInput model.RegisterUser) (*models.User, error) {
	userDB := &models.UserDB{
		Name:        userInput.Name,
		EmailID:     userInput.EmailID,
		PhoneNumber: userInput.PhoneNumber,
		Password:    userInput.Password,
		Address: &models.Address{
			Street:   userInput.Address.Street,
			Landmark: userInput.Address.Landmark,
			City:     userInput.Address.City,
			Country:  userInput.Address.Country,
			ZipCode:  userInput.Address.ZipCode,
		},
	}

	fieldErrors := validateSignupRequest(userDB, a.Validator)

	// TODO: Find good way to show error in response.
	// TODO: There is not concept of status code here so how should we proceed.
	if len(fieldErrors) > 0 {
		return nil, errors.New(fieldErrors[0].Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// save user to database
	result, err := a.Database.Collection(a.DB_Collections.USERS).InsertOne(ctx, userDB)

	if err != nil {
		return nil, err
	}

	userDB.ID = result.InsertedID.(primitive.ObjectID)

	user := &models.User{
		ID:          userDB.ID.Hex(),
		Name:        userDB.Name,
		EmailID:     userDB.EmailID,
		PhoneNumber: userDB.PhoneNumber,
		Address:     userDB.Address,
		AuthToken:   userDB.AuthToken,
		Role:        userDB.Role,
	}

	return user, nil
}
func (a *api) SignInUserByEmail(ctx context.Context, user model.SignInUserByEmail) (*models.User, error) {

	// TODO: Validations for emailId and Password
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// save user to database
	result := a.Database.Collection(a.DB_Collections.USERS).FindOne(ctx, bson.M{
		"emailid": user.EmailID,
	})

	var userDB models.UserDB
	err := result.Decode(&userDB)
	if err != nil {
		return nil, fmt.Errorf("wrong crendetials provided.")
	} else {
		// check if password is correct and found email id
		if userDB.Password != user.Password {
			return nil, errors.New("wrong credentials provided.")
		}

		// generate token
		claims := &models.UserClaims{
			ObjectId: userDB.ID,
			EmailId:  userDB.EmailID,
			Name:     userDB.Name,
			Role:     userDB.Role,
		}

		authToken, err := a.AuthService.GenerateJwtToken(claims)

		if err != nil {
			// internal error handle
			return nil, fmt.Errorf("error generating auth token : %v", err.Error())
		}

		responseUser := &models.User{
			ID:          userDB.ID.Hex(),
			Name:        userDB.Name,
			EmailID:     userDB.EmailID,
			PhoneNumber: userDB.PhoneNumber,
			Role:        userDB.Role,
			Address: &models.Address{
				Street:   userDB.Address.Street,
				Landmark: userDB.Address.Landmark,
				City:     userDB.Address.City,
				Country:  userDB.Address.Country,
				ZipCode:  userDB.Address.ZipCode,
			},
			AuthToken: &authToken,
		}

		return responseUser, nil

	}
}

func (a *api) GetUser(ctx context.Context, authToken string) (*models.User, error) {
	userClaims, err := a.AuthService.VerifyJwt(authToken)

	if err != nil {
		return nil, fmt.Errorf("error verifying token : invalid token : %v", err.Error())
	}

	result := a.Database.Collection(a.DB_Collections.USERS).FindOne(ctx, bson.M{
		"_id": userClaims.ObjectId,
	})

	var userDB models.UserDB

	err = result.Decode(&userDB)
	if err != nil {
		return nil, fmt.Errorf("error decoding user from database : invalid token : %v", err.Error())
	}

	responseUser := &models.User{
		ID:          userDB.ID.Hex(),
		Name:        userDB.Name,
		EmailID:     userDB.EmailID,
		PhoneNumber: userDB.PhoneNumber,
		Role:        userDB.Role,
		Address: &models.Address{
			Street:   userDB.Address.Street,
			Landmark: userDB.Address.Landmark,
			City:     userDB.Address.City,
			Country:  userDB.Address.Country,
			ZipCode:  userDB.Address.ZipCode,
		},
		AuthToken: &authToken,
	}

	return responseUser, nil
}

func (a *api) VerifyUserEmail(ctx context.Context, authToken string) (*models.User, error) {

	userClaims, err := a.AuthService.VerifyJwt(authToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying token : ", err.Error())
	}

	companyName := "Graphql Eccomerce"

	url := fmt.Sprintf("http://localhost:8080/verifyEmail?authToken=%v", authToken)
	htmlContent := fmt.Sprintf(`
    <html>
    <head>
        <title>Verify Your Email Address</title>
    </head>
    <body>
        <h1>Welcome to %s!</h1>
        <p>Hello %s,</p>
        <p>Thank you for registering with us! To ensure the security of your account and to activate all the features associated with it, we kindly ask you to verify your email address.</p>
		<h2> CLICK ON LINK BELOW TO VERIFY YOUR EMAIL </h2>
		<a href="%s">VERIFY EMAIL</a>
    </body>
    </html>`,
		companyName,
		userClaims.Name,
		url,
	)
	// send mail to user for verifying email
	err = a.MailService.SendMail(userClaims.EmailId, htmlContent)

	if err != nil {
		return nil, fmt.Errorf("error sending mail %v", err.Error())
	}
	// use email service

	// find given user and send it.
	result := a.Database.Collection(a.DB_Collections.USERS).FindOne(ctx, bson.M{
		"_id": userClaims.ObjectId,
	})

	var userDB models.UserDB

	err = result.Decode(&userDB)
	if err != nil {
		return nil, fmt.Errorf("error decoding user from database : invalid token : %v", err.Error())
	}

	responseUser := &models.User{
		ID:          userDB.ID.Hex(),
		Name:        userDB.Name,
		EmailID:     userDB.EmailID,
		PhoneNumber: userDB.PhoneNumber,
		Role:        userDB.Role,
		Address: &models.Address{
			Street:   userDB.Address.Street,
			Landmark: userDB.Address.Landmark,
			City:     userDB.Address.City,
			Country:  userDB.Address.Country,
			ZipCode:  userDB.Address.ZipCode,
		},
		AuthToken: &authToken,
	}

	return responseUser, nil
}

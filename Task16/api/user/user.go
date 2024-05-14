package user

import (
	"context"
	"errors"
	"graphql_search/models"
	"time"
)

func (a *api) RegisterUser(ctx context.Context, name string, emailID string, phoneNumber string, address models.AddressInput) (*models.User, error) {
	user := &models.User{
		Name:        name,
		EmailID:     emailID,
		PhoneNumber: phoneNumber,
		Address: &models.Address{
			Street:   address.Street,
			Landmark: address.Landmark,
			City:     address.City,
			Country:  address.City,
			ZipCode:  address.ZipCode,
		},
	}

	fieldErrors := validateSignupRequest(user, a.Validator)

	// TODO: Find good way to show error in response.
	// TODO: There is not concept of status code here so how should we proceed.
	if len(fieldErrors) > 0 {
		// handle errors
		return nil, errors.New("field validation error")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// save user to database
	a.Database.Collection(a.DB_Collections.USERS).InsertOne(ctx, user)

	return user, nil
}
func (a *api) SignInUser(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}
func (a *api) GetAllUsers(ctx context.Context, authToken string) ([]*models.User, error) {
	return nil, nil
}

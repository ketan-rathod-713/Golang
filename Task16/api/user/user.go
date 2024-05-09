package user

import (
	"context"
	"graphql_search/graph/model"
)

func (a *api) RegisterUser(ctx context.Context, name string, emailID string, phoneNumber string, address model.AddressInput) (*model.User, error) {
	return nil, nil
}
func (a *api) SignInUser(ctx context.Context, id string) (*model.User, error) {
	return nil, nil
}
func (a *api) GetAllUsers(ctx context.Context, authToken string) ([]*model.User, error) {
	return nil, nil
}

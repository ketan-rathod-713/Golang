package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.46

import (
	"context"
	"graphql_search/models"
)

// GetAllBoardsByTitle is the resolver for the GetAllBoardsByTitle field.
func (r *queryResolver) GetAllBoardsByTitle(ctx context.Context, title string) ([]*models.Board, error) {
	return r.Api.BoardApi.GetBoardsByTitle(title)
}

// GetBoard is the resolver for the GetBoard field.
func (r *queryResolver) GetBoard(ctx context.Context, id string) (*models.Board, error) {
	return r.Api.BoardApi.GetBoard(id)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
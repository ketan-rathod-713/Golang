package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.46

import (
	"context"
	"graphql_search/models"
)

// Products is the resolver for the Products field.
func (r *categoryResolver) Products(ctx context.Context, obj *models.Category) ([]*models.Product, error) {
	// get all products in given category.
	return r.Api.CategoryApi.GetProductsByCategory(obj)
}

// CreateProduct is the resolver for the CreateProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, name string, description string, price float64, quantity int, category string) (*models.Product, error) {
	return r.Api.ProductApi.Create(name, description, price, quantity, category)
}

// CreateCategory is the resolver for the CreateCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, name string) (*models.Category, error) {
	return r.Api.CategoryApi.Create(name)
}

// Category is the resolver for the Category field.
func (r *productResolver) Category(ctx context.Context, obj *models.Product) (*models.Category, error) {
	return r.Api.CategoryApi.Get(obj.Category.ID)
}

// GetProducts is the resolver for the GetProducts field.
func (r *queryResolver) GetProducts(ctx context.Context, pagination *models.Pagination) ([]*models.Product, error) {
	return r.Api.ProductApi.GetAll(pagination)
}

// GetProduct is the resolver for the GetProduct field.
func (r *queryResolver) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	return r.Api.ProductApi.Get(id)
}

// GetCategories is the resolver for the getCategories field.
func (r *queryResolver) GetCategories(ctx context.Context, pagination *models.Pagination) ([]*models.Category, error) {
	// We don't need pagination for categories.
	return r.Api.CategoryApi.GetAll(nil)
}

// GetCategory is the resolver for the getCategory field.
func (r *queryResolver) GetCategory(ctx context.Context, id string) (*models.Category, error) {
	return r.Api.CategoryApi.Get(id)
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Product returns ProductResolver implementation.
func (r *Resolver) Product() ProductResolver { return &productResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

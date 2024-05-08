package category

import (
	"context"
	"graphql_search/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *api) Create(name string) (*models.Category, error) {
	var category models.Category = models.Category{
		Name: name,
	}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	result, err := a.Database.Collection("category").InsertOne(ctx, bson.M{
		name: name,
	})

	if err != nil {
		return nil, err
	}

	category.ID = result.InsertedID.(primitive.ObjectID).Hex()

	log.Println("Category Inserted ", category.ID)
	return &category, nil
}

func (a *api) Get(id string) (*models.Category, error) {
	return nil, nil
}

// Get all categories.
// if pagination is required then do else fetch all categories.
func (a *api) GetAll(pagination *models.Pagination) ([]*models.Category, error) {
	return nil, nil
}

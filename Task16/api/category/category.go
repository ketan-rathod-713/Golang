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
		"name": name,
	})

	if err != nil {
		return nil, err
	}

	category.ID = result.InsertedID.(primitive.ObjectID).Hex()

	log.Println("Category Inserted ", category.ID)
	return &category, nil
}

func (a *api) Get(id string) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)

	result := a.Database.Collection("category").FindOne(ctx, bson.M{
		"_id": objectId,
	})

	var categoryDB models.CategoryDB
	err := result.Decode(&categoryDB)
	if err != nil {
		return nil, err
	}

	category := &models.Category{
		ID:   categoryDB.ID.Hex(),
		Name: categoryDB.Name,
	}

	return category, nil
}

// Get all categories.
// if pagination is required then do else fetch all categories.
func (a *api) GetAll(pagination *models.Pagination) ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}

	cursor, err := a.Database.Collection("category").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var categoryDB []*models.CategoryDB
	err = cursor.All(ctx, &categoryDB)
	if err != nil {
		return nil, err
	}

	var categories []*models.Category
	for _, p := range categoryDB {
		var category *models.Category = &models.Category{
			ID:   p.ID.Hex(),
			Name: p.Name,
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (a *api) GetProductsByCategory(category *models.Category) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"category": category.ID,
	}

	log.Println("Get Products By Category ", filter)

	cursor, err := a.Database.Collection("products").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var productsDB []*models.ProductDB
	err = cursor.All(ctx, &productsDB)
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	for _, p := range productsDB {
		var product *models.Product = &models.Product{
			ID:          p.Id.Hex(),
			Name:        p.Name,
			Description: p.Description,
			Quantity:    p.Quantity,
			Price:       p.Price,
			Category: &models.Category{
				ID: p.Category,
			},
		}

		products = append(products, product)
	}

	return products, nil
}

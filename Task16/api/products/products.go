package products

import (
	"context"
	"graphql_search/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a api) GetAll(pagination *models.Pagination) ([]*models.Product, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var productsDB []*models.ProductDB

	cursor, err := a.Database.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &productsDB)
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	for _, p := range productsDB {
		products = append(products, &models.Product{
			ID:          p.Id.Hex(),
			Name:        p.Name,
			Description: p.Description,
			Quantity:    p.Quantity,
			Price:       p.Price,
			Category: &models.Category{
				ID: p.Category,
			},
		})
	}

	return products, nil
}

func (a *api) Get() {

}

func (a *api) Create(name string, description string, price float64, quantity int, category string) (*models.Product, error) {
	var product models.Product = models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		Category: &models.Category{ // How to fetch only if category is specified. make resolution for it.
			ID: category,
		},
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// ! why here unable to write price and quantity without doube qoutes.

	result, err := a.Database.Collection("products").InsertOne(ctx, bson.M{
		"name":        name,
		"description": description,
		"price":       price,
		"quantity":    quantity,
		"category":    category,
	})

	if err != nil {
		return nil, err
	}

	product.ID = result.InsertedID.(primitive.ObjectID).Hex()

	log.Println("Product Inserted ", product.ID)

	return &product, nil
}

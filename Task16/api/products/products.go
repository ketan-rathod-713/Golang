package products

import (
	"context"
	"graphql_search/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a api) GetAll(ctx context.Context, pagination *models.Pagination) ([]*models.Product, error) {
	log.Println("Fetch All Products")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var productsDB []*models.ProductDB

	cursor, err := a.Database.Collection(a.DB_Collections.PRODUCTS).Find(ctx, bson.M{})
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

func (a *api) Get(ctx context.Context, id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(id)

	var productDB models.ProductDB
	result := a.Database.Collection(a.DB_Collections.PRODUCTS).FindOne(ctx, bson.M{
		"_id": objectId,
	})

	err := result.Decode(&productDB)
	if err != nil {
		return nil, err
	}

	product := &models.Product{
		ID:          productDB.Id.Hex(),
		Name:        productDB.Name,
		Description: productDB.Description,
		Quantity:    productDB.Quantity,
		Price:       productDB.Price,
		Category: &models.Category{
			ID: productDB.Category,
		},
	}

	return product, nil
}

func (a *api) Create(ctx context.Context, name string, description string, price float64, quantity int, category string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product models.Product = models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		Category: &models.Category{ // How to fetch only if category is specified. make resolution for it.
			ID: category,
		},
	}

	// ! why here unable to write price and quantity without double qoutes.

	result, err := a.Database.Collection(a.DB_Collections.PRODUCTS).InsertOne(ctx, bson.M{
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

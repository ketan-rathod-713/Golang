package products

import (
	"context"
	"fmt"
	"graphql_search/models"
	"graphql_search/service/logs"
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

	// check if given product is in redis if it is then return it.
	productCashed, err := a.RedisProductService.Get(id)

	// if found product then return it.
	if err == nil {
		product := &models.Product{
			ID:          productCashed.Id.Hex(),
			Name:        productCashed.Name,
			Description: productCashed.Description,
			Quantity:    productCashed.Quantity,
			Price:       productCashed.Price,
			Status:      productCashed.Status,
			Category: &models.Category{
				ID: productCashed.Category,
			}}
		return product, nil
	}

	objectId, _ := primitive.ObjectIDFromHex(id)

	var productDB models.ProductDB
	result := a.Database.Collection(a.DB_Collections.PRODUCTS).FindOne(ctx, bson.M{
		"_id": objectId,
	})

	err = result.Decode(&productDB)
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

	// Save data to redis cashe
	err = a.RedisProductService.Set(&productDB)
	if err != nil {
		log.Println("error :", err)
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

	lg := logs.New(a.Database, *a.DB_Collections)
	lg.LogInformationToDB(&logs.LogData{Type: "PRODUCT_CREATED", Information: fmt.Sprintf("PRODUCT ID:%v", product.ID), Prioririty: "COMMON_LOGS", CreatedAt: time.Now().String()})
	if err != nil {
		log.Println("ERROR LOGGING")
	}

	return &product, nil
}

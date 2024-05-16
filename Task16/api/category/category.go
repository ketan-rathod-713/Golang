package category

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

func (a *api) Create(ctx context.Context, name string) (*models.Category, error) {
	var category models.Category = models.Category{
		Name: name,
	}

	result, err := a.Database.Collection(a.DB_Collections.CATEGORY).InsertOne(ctx, bson.M{
		"name": name,
	})

	if err != nil {
		return nil, err
	}

	category.ID = result.InsertedID.(primitive.ObjectID).Hex()

	log.Println("Category Inserted ", category.ID)

	lg := logs.New(a.Database, *a.DB_Collections)
	err = lg.LogInformationToDB(&logs.LogData{Type: "CATEGORY_CREATED", Information: fmt.Sprintf("CATEGORY ID:%v", category.ID), Prioririty: "COMMON_LOGS", CreatedAt: time.Now().String()})
	if err != nil {
		log.Println("ERROR LOGGING")
	}

	return &category, nil
}

func (a *api) Get(ctx context.Context, id string) (*models.Category, error) {

	// check if category stored in cashe
	categoryDB, err := a.RedisCategoryService.Get(id)
	if err == nil {
		category := &models.Category{
			ID:   categoryDB.ID.Hex(),
			Name: categoryDB.Name,
		}

		return category, nil
	}

	loader := a.GetCategoryLoader(ctx)
	cb, err := loader.Load(id)

	if err != nil {
		return nil, err
	}

	category := &models.Category{
		ID:   cb.ID.Hex(),
		Name: cb.Name,
	}

	// Save data to redis cashe
	err = a.RedisCategoryService.Set(&cb)
	if err != nil {
		log.Println("error :", err)
	}

	return category, nil
}

// Get all categories.
// if pagination is required then do else fetch all categories.
func (a *api) GetAll(ctx context.Context, pagination *models.Pagination) ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}

	cursor, err := a.Database.Collection(a.DB_Collections.CATEGORY).Find(ctx, filter)
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

func (a *api) GetProductsByCategory(ctx context.Context, category *models.Category) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"category": category.ID,
	}

	log.Println("Get Products By Category ", filter)

	cursor, err := a.Database.Collection(a.DB_Collections.PRODUCTS).Find(ctx, filter)
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

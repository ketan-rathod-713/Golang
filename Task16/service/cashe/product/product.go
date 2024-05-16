package product

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"graphql_search/models"
	"time"
)

const resetTime = 100 * time.Second

func (s *service) Get(id string) (*models.ProductDB, error) {
	productKey := fmt.Sprintf("product:%v", id)
	productData, err := s.Client.Get(context.Background(), productKey).Result()
	if err != nil {
		return nil, errors.New("value not found in redis cashe.")
	}
	// get product data and return it before making call to database
	var productDB *models.ProductDB
	err = json.Unmarshal([]byte(productData), &productDB)

	if err != nil {
		return nil, err
	}

	return productDB, nil
}

func (s *service) Set(productDb *models.ProductDB) error {
	productKey := fmt.Sprintf("product:%v", productDb.Id.Hex())

	productStr, err := json.Marshal(productDb)
	if err == nil {
		_, err := s.Client.Set(context.Background(), productKey, productStr, resetTime).Result()

		if err != nil {
			return err
		}
	}

	return nil
}

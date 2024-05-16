package category

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"graphql_search/models"
	"time"
)

const resetTime = 100 * time.Second

func (s *service) Get(id string) (*models.CategoryDB, error) {
	key := fmt.Sprintf("category:%v", id)
	categoryData, err := s.Client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, errors.New("value not found in redis cashe.")
	}
	// get product data and return it before making call to database
	var categoryDB *models.CategoryDB
	err = json.Unmarshal([]byte(categoryData), &categoryDB)

	if err != nil {
		return nil, err
	}

	return categoryDB, nil
}

func (s *service) Set(categoryDB *models.CategoryDB) error {
	key := fmt.Sprintf("category:%v", categoryDB.ID.Hex())

	str, err := json.Marshal(categoryDB)
	if err == nil {
		_, err := s.Client.Set(context.Background(), key, str, resetTime).Result()

		if err != nil {
			return err
		}
	}

	return nil
}

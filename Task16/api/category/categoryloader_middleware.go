package category

import (
	"context"
	ers "errors"
	"graphql_search/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *api) CategoryLoaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		config := CategoryLoaderConfig{
			Wait:     2 * time.Millisecond,
			MaxBatch: 100,
			Fetch: func(keys []string) ([]models.CategoryDB, []error) {
				// log.Println("Fetch Categories called for keys", keys)

				// data we need to populate.
				categories := make([]models.CategoryDB, len(keys))
				errors := make([]error, len(keys))
				keysObjectIds := make([]primitive.ObjectID, len(keys))

				for i, key := range keys {
					var err error
					keysObjectIds[i], err = primitive.ObjectIDFromHex(key)
					errors[i] = err
				}

				// do batch request here to mongodb
				// Prepare the filter to find documents by their IDs
				filter := bson.M{"_id": bson.M{"$in": keysObjectIds}}
				cursor, err := a.Database.Collection(a.DB_Collections.CATEGORY).Find(context.Background(), filter)
				if err != nil {
					// Handle the error
					for i := range errors {
						errors[i] = err
					}
					return categories, errors
				}
				defer cursor.Close(context.Background())

				// Iterate over the cursor and populate the categories slice
				var categoriesGot []models.CategoryDB
				err = cursor.All(context.Background(), &categoriesGot)
				if err != nil {
					// Handle the error
					for i := range errors {
						errors[i] = err
					}
					return categories, errors
				}

				// log.Println("Categories got for given keys ", categoriesGot)

				// iterate categories got and return exact sequence of data we got in keys
				var mp map[string]models.CategoryDB = make(map[string]models.CategoryDB)
				for _, d := range categoriesGot {
					mp[d.ID.Hex()] = models.CategoryDB{
						ID:   d.ID,
						Name: d.Name,
					}
				}

				// now iterate all keys and enter relevent information
				for i, key := range keys {
					val, ok := mp[key]
					if !ok {
						errors[i] = ers.New("An Error getting data from map.")
					} else {
						categories[i] = val
					}
				}

				return categories, errors
			},
		}

		categoryLoader := NewCategoryLoader(config)

		// add the categoryloader inside the context.

		ctx := context.WithValue(r.Context(), a.CategoryLoaderKey, categoryLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *api) GetCategoryLoader(ctx context.Context) *CategoryLoader {
	return ctx.Value(a.CategoryLoaderKey).(*CategoryLoader)
}

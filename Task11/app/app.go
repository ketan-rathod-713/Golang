package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodbpagination/models"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to database with timeout of 20 sec.
func InitDB(env *models.Env) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.DB_URL))

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// get database for performing operations
func GetDB(client *mongo.Client) *mongo.Database {
	return client.Database("task10_mongodb_train_dataset")
}

func GetCollectionTrain(db *mongo.Database) *mongo.Collection {
	return db.Collection("trains")
}

// load env variables
func LoadEnv() *models.Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := models.Env{
		PORT:   os.Getenv("PORT"),
		DB_URL: os.Getenv("DB_URL"),
	}

	return &env
}

// don't use special characters in regex and hence remove it before.
func escapeRegex(text string) string {
	return regexp.QuoteMeta(text)
}

func StartServer() {
	// load env variables
	env := LoadEnv()

	// intialise database
	client, err := InitDB(env)
	if err != nil {
		log.Fatal(err)
	}

	db := GetDB(client)
	trainsColl := GetCollectionTrain(db)

	// API STATUS
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// GET /trains 
	// Example URL : /trains?page=1&limit=10&search=ahmedabad&sort="name"&order=1
	// Get Paginated trains data with sorting of one field.
	// If query params not given then default values will be as follows :
	// page : 1, limit : 10, sort : "number", order : 1, search : ""
	// TODO not using controllers for now. 
	http.HandleFunc("/train", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodGet {
			queries, err := url.ParseQuery(r.URL.RawQuery)
			if err != nil {
				json.NewEncoder(w).Encode(models.ErrorResponse{Status: 404, Message: "Invalid query"})
				log.Println(err)
				return
			}

			var page int = 1
			var limit int = 10
			var sortField string = "number"
			var sortOrder int = 1
			var searchText string = ""

			// get returns the first query
			pageStr := queries.Get("page")
			limitStr := queries.Get("limit")
			sortFieldStr := queries.Get("sort")
			sortOrderStr := queries.Get("order")
			searchTextStr := queries.Get("search")

			if searchTextStr != "" {
				searchText = escapeRegex(searchTextStr)
			}

			log.Println(searchTextStr)

			if sortOrderStr != "" {
				sortOrder, err = strconv.Atoi(sortOrderStr)
				if err != nil {
					json.NewEncoder(w).Encode(models.ErrorResponse{Status: 404, Message: "Sort Order String to Int conversion failed"})
					log.Println(err)
					return
				}
			}

			if pageStr != "" {
				page, err = strconv.Atoi(pageStr)
				if err != nil {
					json.NewEncoder(w).Encode(models.ErrorResponse{Status: 404, Message: "Page Number string to int conversion failed"})
					log.Println(err)
					return
				}
			}

			if limitStr != "" {
				limit, err = strconv.Atoi(limitStr)
				if err != nil {
					json.NewEncoder(w).Encode(models.ErrorResponse{Status: 404, Message: "Limit number string to int conversion failed"})
					log.Println(err)
					return
				}
			}

			trainNumber, err := strconv.Atoi(searchText)
			if err != nil {
				trainNumber = 0
			}

			filter := bson.M{
				"$or": []bson.M{
					bson.M{"number": trainNumber}, // rather then directly matching i want regex here
					bson.M{"name": primitive.Regex{Pattern: "^" + searchText + ".*", Options: "i"}},
					bson.M{"source": primitive.Regex{Pattern: "^" + searchText + ".*", Options: "i"}},
					bson.M{"destination": primitive.Regex{Pattern: "^" + searchText + ".*", Options: "i"}},
				},
			}

			if sortFieldStr != "" {
				sortField = sortFieldStr
			} else {
				filter = bson.M{}
			}

			// Now fetch data from mongodb
			var skip = (page - 1) * limit
			cursor, err := trainsColl.Find(context.Background(), filter, options.Find().SetLimit(int64(limit)).SetSkip(int64(skip)).SetSort(bson.M{sortField: sortOrder}))
			if err != nil {
				json.NewEncoder(w).Encode(models.ErrorResponse{Status: 404, Message: "error finding data from trains collection"})
				log.Println(err)
				return
			}

			var trains []models.Train

			for cursor.Next(context.Background()) {
				var train models.Train
				err := cursor.Decode(&train)
				if err != nil {
					json.NewEncoder(w).Encode(models.ErrorResponse{Status: 404, Message: "error decoding data from trains collection"})
					log.Println(err)
					return
				}
				trains = append(trains, train)
			}

			// ? Below way is also valid but let me try above one.
			// err = cursor.All(context.Background(), &trains)
			// if err != nil {
			// 	json.NewEncoder(w).Encode(models.ErrorResponse{Status: 404, Message: "Error gettting data from cursor"})
			// 	log.Println(err)
			// 	return
			// }

			totalDocuments, err := trainsColl.CountDocuments(context.Background(), filter)
			if err != nil {
				json.NewEncoder(w).Encode(models.ErrorResponse{Status: 404, Message: "Error getting total count of all documents"})
				log.Println(err)
				return
			}

			if len(trains) == 0 {
				json.NewEncoder(w).Encode(models.ErrorResponse{Status: 200, Message: "No trains found with given query"})
				log.Println(trains)
				return
			}

			json.NewEncoder(w).Encode(
				models.PaginatedTrainsResponse{
					Page:  page,
					Limit: limit,
					Total: int(totalDocuments),
					Data:  trains,
				},
			)
		}
	})



	// Start Server On Given Port
	log.Println("Server started on port", env.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", env.PORT), nil))
}

package books

import (
	"books/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type getAllBooksResponse struct {
	Books []models.Book `json:"books"`
}

// get all books without authorization
func (a *Api) HandleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	cursor, err := a.App.DB.Collection("books").Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error while getting books from database", err)

		return
	}

	fmt.Println(cursor.Current.Elements())

	var books []models.Book
	err = cursor.All(context.TODO(), &books)
	if err != nil {
		log.Println("Error while getting books from database", err)

		return
	}

	json.NewEncoder(w).Encode(getAllBooksResponse{
		Books: books,
	})
}

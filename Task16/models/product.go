package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          string    `json:"Id" bson:"_id"`
	Name        string    `json:"Name" bson:"name"`
	Description string    `json:"Description" bson:"description"`
	Quantity    int       `json:"Quantity" bson:"quantity"`
	Price       float64   `json:"Price" bson:"price"`
	Category    *Category `json:"Category" bson:"-"`
	Status      string    `json:"Status", bson:"status"`
}

// for db related data transmission
type ProductDB struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string
	Description string
	Quantity    int
	Price       float64
	Category    string
	Status      string
}

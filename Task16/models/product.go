package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          string    `json:"Id"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Quantity    int       `json:"Quantity"`
	Price       float64   `json:"Price"`
	Category    *Category `json:"Category" bson:"-"`
}

// for db related data transmission
type ProductDB struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string
	Description string
	Quantity    int
	Price       float64
	Category    string
}

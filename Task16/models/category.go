package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID       string     `json:"Id" bson:"_id"`
	Name     string     `json:"Name"`
	Products []*Product `json:"Products" bson:"-"`
}

// database category model
type CategoryDB struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

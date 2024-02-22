package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // ingore empty fields and make this field required
	Name     string             `json:"name,omitempty" validate:"required" bson:"name"`
	Location string             `json:"location,omitempty" validate:"required" bson:"location"`
	Title    string             `json:"title,omitempty" validate:"required" bson:"title"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Env struct {
	PORT string
}

type Train struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Number      int                `json:"number"`
	Source      string             `json:"source"`
	Destination string             `json:"destination"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // ingore empty fields and make this field required
	Name          string             `json:"name,omitempty" validate:"required" bson:"name"`
	Location      string             `json:"location,omitempty" validate:"required" bson:"location"`
	Title         string             `json:"title,omitempty" validate:"required" bson:"title"`
	Age           int64              `json:"age,omitempty" validate:"required" bson:"age"`
	FavoriteGames []string           `json:"favoriteGames,omitempty" validate:"required" bson:"favoriteGames"`
	Hobby         Hobby              `json:"hobby,omitempty" validate:"required" bson:"hobby"`
}

type Hobby struct {
	Name  string `json:"name,omitempty" validate:"required" bson:"name"`
	Years int    `json:"years,omitempty" validate:"required" bson:"years"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Config struct {
	DB_URL    string
	REST_PORT string
	SECRET    string
	DATABASE  string
	GRPC_PORT string
}

type User struct {
	ObjectId *primitive.ObjectID `json:"object_id" bson:"_id,omitempty"`
	Email    string              `json:"email" validate:"required" bson:"email,omitempty"`
	Password string              `json:"-" validate:"required" bson:"password,omitempty"`
	Name     string              `json:"name" validate:"required" bson:"name,omitempty"`
	Phone    string              `json:"phone,omitempty" bson:"phone,omitempty"`
	Address  string              `json:"address,omitempty" bson:"address,omitempty"`
	City     string              `json:"city,omitempty" bson:"city,omitempty"`
	State    string              `json:"state,omitempty" bson:"state",omitempty"`
	Country  string              `json:"country,omitempty" bson:"country,omitempty"`
	Zip      string              `json:"zip,omitempty" bson:"zip,omitempty"`
	Standard string              `json:"standard" validate:"required" bson:"standard,omitempty"`
	JwtToken string              `json:"jwt_token" bson:"-"`
	Role     string              `json:"role" bson:"role,omitempty" validate:"required"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Config struct {
	DB_URL    string
	REST_PORT string
	SECRET    string
	DATABASE  string
	GRPC_PORT string
}

// actually stored in database with qty and all fields, all fields can only be view by admin only
type Book struct {
	ObjectId  *primitive.ObjectID `json:"object_id" bson:"_id,omitempty"`
	CreatedBy *primitive.ObjectID `json:"created_by" bson:"created_by,omitempty"`
	Title     string              `json:"title" validate:"required" bson:"title,omitempty"`
	Author    string              `json:"author" validate:"required" bson:"author,omitempty"`
	Qty       int                 `json:"qty" validate:"required" bson:"qty,omitempty"`
	Price     float64             `json:"price" validate:"required" bson:"price,omitempty"`
	Image     string              `json:"image" validate:"required" bson:"image,omitempty"`
	Category  string              `json:"category" validate:"required" bson:"category,omitempty"` // Action, Adventure etc.
	Status    string              `json:"status" validate:"required" bson:"status,omitempty"`
	Created   string              `json:"created" validate:"required" bson:"created,omitempty"`
	Updated   string              `json:"updated" validate:"required" bson:"updated,omitempty"`
	IsDeleted bool                `json:"isDeleted" bson:"isDeleted"` // Store Deleted Books Too
}

// how to declare enums in golang

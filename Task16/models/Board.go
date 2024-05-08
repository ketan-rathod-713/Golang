package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Board struct {
	ID          string `json:"Id"`
	BoardID     string `json:"boardId"`
	Visible     string `json:"visible"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type BoardDB struct {
	ID          primitive.ObjectID `json:"Id" bson:"_id"`
	BoardID     BoardIdData        `json:"boardId"`
	Visible     string             `json:"visible"`
	Description string             `json:"description"`
	Title       string             `json:"title"`
	Type        string             `json:"type"`
}

type BoardIdData struct {
	Id primitive.ObjectID `bson:"_id"`
}

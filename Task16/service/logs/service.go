package logs

import (
	"graphql_search/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogData struct {
	Id          primitive.ObjectID `bson:"_id"`
	Type        string             `bson:"type"`
	Information string             `bson:"information"`
	Prioririty  string             `bson:"priority"`
	CreatedAt   string             `bson:"createdAt"`
}

// all jwt auth services define here
type service struct {
	DB            *mongo.Database
	DBCollections *models.DB_COLLECTIONS
}

func New(db *mongo.Database, dbCollections models.DB_COLLECTIONS) Service {
	return &service{
		DB:            db,
		DBCollections: &dbCollections,
	}
}

type Service interface {
	LogInformationToDB(logdata *LogData) error
	UpdateTTLIndex(newTime time.Duration) error
}

package logs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) LogInformationToDB(logdata *LogData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log_data := bson.M{
		"type":        logdata.Type,
		"information": logdata.Information,
		"priority":    logdata.Prioririty,
		"createdAt":   logdata.CreatedAt,
	}

	result, err := s.DB.Collection(s.DBCollections.LOGS).InsertOne(ctx, log_data)

	if err != nil {
		return err
	}

	logdata.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (s *service) UpdateTTLIndex(newTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	raw, err := s.DB.Collection(s.DBCollections.LOGS).Indexes().DropOne(ctx, "expiry")
	if err != nil {
		return err
	}

	log.Println("Index Deleted ", raw)

	// Create new index
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"expiry": 1},                                             // Index key, ascending order
		Options: options.Index().SetExpireAfterSeconds(int32(newTime.Seconds())), // Automatically delete documents after specified seconds
	}
	_, err = s.DB.Collection(s.DBCollections.LOGS).Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	log.Println("Index Created Successfully")
	return nil
}

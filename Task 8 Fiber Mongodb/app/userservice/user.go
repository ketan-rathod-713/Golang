package userservice

import (
	"context"
	"errors"
	"fibermongoapp/models"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) CreateUser(user *models.User) (*models.User, error) {

	result, err := s.UserCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	log.Info("Inserted User Id Is ", result.InsertedID)

	// TODO Add Type Assertion Here
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}

	return user, nil
}

func (s *service) GetUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.UserCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var users []*models.User
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// require id params
func (s *service) GetOneUserById(userId primitive.ObjectID) (*models.User, error) {

	var user *models.User
	result := s.UserCollection.FindOne(context.TODO(), bson.M{"_id": userId})
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// require id params
func (s *service) UpdateUser(user *models.User) (*models.User, error) {

	update := bson.D{
		{"$set", user},
	}

	result, err := s.UserCollection.UpdateByID(context.TODO(), user.ID, update, options.Update().SetUpsert(false))
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("No documents matching given filters")
	}

	if result.ModifiedCount == 1 {
		return user, nil
	}

	return nil, errors.New("No documents modiefied")
}

// require id params
func (s *service) DeleteUser(userId primitive.ObjectID) (*models.User, error) {

	// TODO start Transaction Here
	// ! Need to start transaction from here
	var user *models.User
	userResult := s.UserCollection.FindOne(context.TODO(), bson.M{"_id": userId})
	err := userResult.Decode(&user)
	if err != nil {
		return nil, err
	}

	result, err := s.UserCollection.DeleteOne(context.TODO(), bson.D{{"_id", userId}})

	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, errors.New("No Records Deleted")
	}

	return user, nil
}

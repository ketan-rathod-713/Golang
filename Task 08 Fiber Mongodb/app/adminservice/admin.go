package adminservice

import (
	"context"
	"fibermongoapp/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) CreateSubject(subject *models.Subject) error {
	result, err := s.App.Collections.Subject.InsertOne(context.TODO(), bson.M{
		"name":     subject.Name,
		"teachers": []primitive.ObjectID{},
	})
	subject.Id = result.InsertedID.(primitive.ObjectID)
	return err
}

func (s *service) CreateClass(class *models.Class) error {
	result, err := s.App.Collections.Class.InsertOne(context.TODO(), bson.M{
		"name":     class.Name,
		"subjects": []primitive.ObjectID{}, // empty array of subjects
	})
	class.Id = result.InsertedID.(primitive.ObjectID)
	return err
}

func (s *service) CreateTeacher(teacher *models.Teacher) error {
	result, err := s.App.Collections.Teacher.InsertOne(context.TODO(), bson.M{
		"name": teacher.Name,
	})
	teacher.Id = result.InsertedID.(primitive.ObjectID)
	return err
}

func (s *service) AddSubjectToClass(classId primitive.ObjectID, subjectId primitive.ObjectID) error {

	filter := bson.M{
		"_id": classId,
	}

	// IMP: addToSet for adding only unique entries in array of document
	update := bson.M{
		"$addToSet": bson.M{
			"subjects": subjectId,
		},
	}
	result := s.App.Collections.Class.FindOneAndUpdate(context.TODO(), filter, update)

	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (s *service) AddTeacherToSubject(teacherId primitive.ObjectID, subjectId primitive.ObjectID) error {

	filter := bson.M{
		"_id": bson.M{
			"_id": subjectId,
		},
	}

	// IMP: addToSet for adding only unique entries in array of document
	update := bson.M{
		"$addToSet": bson.M{
			"teachers": teacherId,
		},
	}

	result := s.App.Collections.Class.FindOneAndUpdate(context.TODO(), filter, update)

	if result.Err() != nil {
		fmt.Println(result.Err())
		return result.Err()
	}

	return nil
}

func (s *service) DeleteSubjectFromClass(classId primitive.ObjectID, subjectId primitive.ObjectID) error {

	filter := bson.M{
		"_id": classId,
	}

	// IMP: pull for removing elements from an array matching given condition
	update := bson.M{
		"$pull": bson.M{
			"subjects": bson.M{
				"$eq": subjectId,
			},
		},
	}
	result := s.App.Collections.Class.FindOneAndUpdate(context.TODO(), filter, update)

	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (s *service) GetClass() ([]*models.Class, error) {
	var classes []*models.Class
	filter := bson.M{}

	aggregate := bson.A{
		bson.M{
			"$match": filter,
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "subjects",
				"localField":   "subjects",
				"foreignField": "_id",
				"as":           "subjects",
			},
		},
	}

	cursor, err := s.App.Collections.Class.Aggregate(context.TODO(), aggregate)

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &classes)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

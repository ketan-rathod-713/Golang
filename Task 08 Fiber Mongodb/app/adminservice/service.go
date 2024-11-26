package adminservice

import (
	"fibermongoapp/app"
	"fibermongoapp/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service interface will serve as a contract for user services
type Service interface {
	CreateSubject(subject *models.Subject) error
	CreateClass(subject *models.Class) error
	CreateTeacher(teacher *models.Teacher) error
	AddSubjectToClass(classId primitive.ObjectID, subjectId primitive.ObjectID) error
	AddTeacherToSubject(teacherId primitive.ObjectID, subjectId primitive.ObjectID) error
	DeleteSubjectFromClass(classId primitive.ObjectID, subjectId primitive.ObjectID) error
	GetClass() ([]*models.Class, error)
}

/*
user service requires db pointer and mongo.Collection pointer to query user collection
TODO: here collection name is hardcoded -> improve it.
*/
type service struct {
	App *app.App
}

func New(a *app.App) Service {
	return &service{
		App: a,
	}
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type StudentRegistration struct {
	Id        primitive.ObjectID `json:"id,omitempty" validate:"required" bson:"_id,omitempty"`
	StudentId primitive.ObjectID `json:"student,omitempty" validate:"required" bson:"student"`
	Year      int                `json:"year,omitempty" validate:"required" bson:"year"`
	ClassId   primitive.ObjectID `json:"class,omitempty" validate:"required" bson:"class"`
	Timestamp string             `json:"timeStamp,omitempty" validate:"required bson:"timestamp"`
}

type Class struct {
	Id       primitive.ObjectID `json:"id,omitempty" validate:"required" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required" bson:"name"`
	Subjects []Subject          `json:"subjects,omitempty" validate:"required" bson:"subjects"`
}

type Subject struct {
	Id       primitive.ObjectID `json:"id,omitempty" validate:"required" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required" bson:"name"`
	Teachers []Teacher          `json:"teacher,omitempty" validate:"required" bson:"teacher"`
}

type Teacher struct {
	Id   primitive.ObjectID `json:"id,omitempty" validate:"required" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" validate:"required" bson:"name"`
}

// TODO add something interesting ??
// let's say user will have class, and class will have subjects

/*
user {
	infor about him
}

class { // accoring to top levels
	id
	name
	subjects: array of subjects // add subject and all that queries
}

subject {
	id
	name
}

student_registrations {
	id
	student id NOT UNIQUE
	class id
	year // for which year registered
	time_stamp // time at which registered.
	? any other info
}

student_class {
	studentId and classId
}
*/

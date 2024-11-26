package graph

import (
	"meetmeup/models"
	"meetmeup/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var meetups = []*models.Meetup{
	&models.Meetup{
		ID:          "1",
		Name:        "a meetup",
		Description: "description",
		UserId:      "1",
	},
	&models.Meetup{
		ID:          "2",
		Name:        "a meetup",
		Description: "description",
		UserId:      "1",
	},
	&models.Meetup{
		ID:          "3",
		Name:        "a meetup",
		Description: "description",
		UserId:      "2",
	},
}

var users = []*models.User{
	&models.User{
		ID:       "1",
		Username: "ketan",
		Email:    "ketanrtd1@gmail.com",
	},
	&models.User{
		ID:       "2",
		Username: "ketan2",
		Email:    "ketanrtd2@gmail.com",
	},
}

type Resolver struct {
	MeetupRepo postgres.MeetupRepo
	UserRepo   postgres.UserRepo
}

package graph

import (
	"context"
	"errors"
	"fmt"
	"meetmeup/graph/model"
	"meetmeup/models"
)

type mutationResolver struct{ *Resolver }

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	var user models.User = models.User{
		Username: input.Username,
		Email:    input.Email,
	}
	return r.UserRepo.CreateUser(user)
}

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {

	// do some validations

	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserId:      input.UserID,
	}

	return r.MeetupRepo.CreateMeetup(meetup)
}

// UpdateMeetup is the resolver for the updateMeetup field.
func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*models.Meetup, error) {
	panic(fmt.Errorf("not implemented: UpdateMeetup - updateMeetup"))
}

package graph

import (
	"context"
	"errors"
	"meetmeup/graph/model"
	"meetmeup/models"
)

type mutationResolver struct{ *Resolver }

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
		UserId:      "1",
	}

	return r.MeetupRepo.CreateMeetup(meetup)
}

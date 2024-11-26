package graph

import (
	"context"
	"meetmeup/graph/model"
	"meetmeup/models"
)

type queryResolver struct{ *Resolver }

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context, input *model.Search) ([]*models.Meetup, error) {

	if input == nil {
		// logic if didn't get input
		return r.MeetupRepo.GetMeetups()
	} else {
		if input.SearchText == nil {
			return r.MeetupRepo.GetMeetups()
		} else {
			// logic if got search text
			return r.MeetupRepo.GetMeetupsBySearchText(*input.SearchText)
		}
	}
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	// fetch all the users and return it.
	return r.UserRepo.GetAllUsers()
}

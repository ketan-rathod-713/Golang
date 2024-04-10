package graph

import (
	"context"
	"meetmeup/models"
)

type queryResolver struct{ *Resolver }

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupRepo.GetMeetups()
}


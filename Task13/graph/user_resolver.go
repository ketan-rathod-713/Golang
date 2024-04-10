package graph

import (
	"context"
	"meetmeup/models"
)

type userResolver struct{ *Resolver }

// User returns UserResolver implementation.


// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {

	// ! Without using databse
	// // all meetups for current user
	// result := make([]*models.Meetup, 0)

	// for _, m := range meetups {
	// 	if m.UserId == obj.ID {
	// 		result = append(result, m)
	// 	}
	// }

	// ? Using database

	return nil, nil
}

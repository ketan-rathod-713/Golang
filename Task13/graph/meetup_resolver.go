package graph

import (
	"context"
	"fmt"
	"meetmeup/models"
)

type meetupResolver struct{ *Resolver }

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {

	// ! Without using database
	// user := new(models.User)
	// // loop on all users and get user
	// for _, u := range users {
	// 	if u.ID == obj.UserId {
	// 		user = u
	// 		break
	// 	}
	// }

	// if user == nil {
	// 	return nil, errors.New("user with id not exist")
	// }

	// ? Using database
	// return r.UserRepo.GetUserById(obj.UserId)

	// This query is coming from dataloader middleware already.

	// ? with data loader
	// TODO understand how it works internally.

	// kya loader k pas sare users start me hi aa jaege. or it will just cashe the data whatever it gets.
	data, err := GetUserLoader(ctx).Load(obj.UserId)

	fmt.Println("data", data)
	fmt.Println("err", err)
	return data, err
}

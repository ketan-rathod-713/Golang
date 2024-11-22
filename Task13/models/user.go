package models

type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	// Meetups  []*Meetup `json:"meetups"`
	// we will not get meetups from here, we will get it from the resolver ha ha
}

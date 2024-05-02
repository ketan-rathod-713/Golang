// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewMeetup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Query struct {
}

type UpdateMeetup struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

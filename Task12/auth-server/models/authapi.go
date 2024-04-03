package models

type SignupRequest struct {
	Email    string `json:"email" validate:"required" bson:"email,omitempty"`
	Password string `json:"password" validate:"required" bson:"password,omitempty"`
	Name     string `json:"name" validate:"required" bson:"name,omitempty"`
	Phone    string `json:"phone" bson:"phone,omitempty"`
	Address  string `json:"address" bson:"address,omitempty"`
	City     string `json:"city" bson:"city,omitempty"`
	State    string `json:"state" bson:"state",omitempty"`
	Country  string `json:"country" bson:"country,omitempty"`
	Zip      string `json:"zip" bson:"zip,omitempty"`
	Standard string `json:"standard" validate:"required" bson:"standard,omitempty"`
	Role     string `json:"role" bson:"role,omitempty" validate:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
	Type   string  `json:"type"`
}

type Error struct { // can be validation error, internal error etc.
	Field string `json:"field"`
	Error string `json:"error"`
}

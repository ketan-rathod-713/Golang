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

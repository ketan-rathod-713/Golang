package models

type User struct {
	ID          string   `json:"ID" bson:"id"`
	Name        string   `json:"name" validate:"required" bson:"name"`
	EmailID     string   `json:"emailId" validate:"required,email" bson:"emailid"`
	PhoneNumber string   `json:"phoneNumber" validate:"required" bson:"phonenumber"`
	Address     *Address `json:"address" validate:"required" bson:"address"`
	AuthToken   *string  `json:"authToken,omitempty" bson:"authtoken"`
	Role        string   `json:"role" bson:"role"`
}
package models

type Address struct {
	Street   string `json:"street" validate:"required" bson:"street"`
	Landmark string `json:"landmark" validate:"required" bson:"landmark"`
	City     string `json:"city" validate:"required" bson:"city"`
	Country  string `json:"country" validate:"required" bson:"country"`
	ZipCode  string `json:"zipCode" validate:"required" bson:"zipcode"`
}

type AddressInput struct {
	Street   string `json:"street"`
	Landmark string `json:"landmark"`
	City     string `json:"city"`
	Country  string `json:"country"`
	ZipCode  string `json:"zipCode"`
}
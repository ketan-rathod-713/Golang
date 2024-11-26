package models

type ErrorResponse struct {
	Errors []Error `json:"errors"`
	Type   string  `json:"type"`
}

type Error struct { // can be validation error, internal error etc.
	Field string `json:"field"`
	Error string `json:"error"`
}

type BookCreateRequest struct {
	Title    string  `json:"title" validate:"required" bson:"title,omitempty"`
	Author   string  `json:"author" validate:"required" bson:"author,omitempty"`
	Qty      int     `json:"qty" validate:"required" bson:"qty,omitempty"`
	Price    float64 `json:"price" validate:"required" bson:"price,omitempty"`
	Category string  `json:"category" validate:"required" bson:"category,omitempty"` // Action, Adventure etc.
	Status   string  `json:"status" validate:"required" bson:"status,omitempty"`
}

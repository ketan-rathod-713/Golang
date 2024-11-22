package jwtauth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claims - a struct that will be encoded to JWT
type UserClaims struct {
	ObjectId primitive.ObjectID `json:"object_id"`
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	Role     string             `json:"role"`
	jwt.RegisteredClaims
}

// JWTToken - JWT Token
type JWTToken struct {
	Value     string
	ExpiresAt time.Time
}

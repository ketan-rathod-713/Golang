package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claims - a struct that will be encoded to JWT
type UserClaims struct {
	ObjectId primitive.ObjectID `json:"objectId"`
	EmailId  string             `json:"email"`
	Name     string             `json:"name"`
	Role     string             `json:"role"`
	jwt.RegisteredClaims
}

type JWTToken struct {
	Value     string
	ExpiresAt time.Time
}

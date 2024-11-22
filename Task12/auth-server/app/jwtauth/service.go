package jwtauth

import (
	"auth/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateJwtToken(user *UserClaims, config *models.Config) (string, error)
	VerifyJwt(jwtToken string, config *models.Config) (*UserClaims, error)
}

// all jwt auth services define here
type service struct {
}

func New() Service {
	return &service{}
}

// generate jwt token for user
func (s *service) GenerateJwtToken(user *UserClaims, config *models.Config) (string, error) {
	// map claims are the extra data that we want to store
	fmt.Println("Generating jwt token for object id", user.ObjectId)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		ObjectId: user.ObjectId,
		Email:    user.Email,
		Name:     user.Name,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 5)},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	})

	var secretByte = []byte(config.SECRET)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := jwtToken.SignedString(secretByte)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// verify jwt and return its claims and error if any
func (s *service) VerifyJwt(jwtToken string, config *models.Config) (*UserClaims, error) {

	var claims UserClaims
	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("unknown claims type, cannot proceed")
	}
}

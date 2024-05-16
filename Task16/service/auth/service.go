package auth

import (
	"graphql_search/models"
	"time"
)

type Service interface {
	GenerateJwtToken(user *models.UserClaims) (string, error)
	VerifyJwt(jwtToken string) (*models.UserClaims, error)
}

// all jwt auth services define here
type service struct {
	ExpireTime time.Duration
	JwtSecret  string
}

func New(jwtSecret string, expireTime time.Duration) Service {
	return &service{
		ExpireTime: expireTime,
		JwtSecret:  jwtSecret,
	}
}

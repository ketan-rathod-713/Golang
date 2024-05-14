package auth

import "graphql_search/models"

type Service interface {
	GenerateJwtToken(user *UserClaims, config *models.Configs) (string, error)
	VerifyJwt(jwtToken string, config *models.Configs) (*UserClaims, error)
}

// all jwt auth services define here
type service struct {
}

func New() Service {
	return &service{}
}

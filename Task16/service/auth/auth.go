package auth

import (
	"errors"
	"graphql_search/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// generate jwt token for user
func (s *service) GenerateJwtToken(user *models.UserClaims) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.UserClaims{
		ObjectId: user.ObjectId,
		EmailId:  user.EmailId,
		Name:     user.Name,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(s.ExpireTime)},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	})

	var secretByte = []byte(s.JwtSecret)

	tokenString, err := jwtToken.SignedString(secretByte)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// verify jwt and return its claims and error if any
func (s *service) VerifyJwt(jwtToken string) (*models.UserClaims, error) {

	var claims models.UserClaims
	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*models.UserClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("unknown claims type, cannot proceed")
	}
}

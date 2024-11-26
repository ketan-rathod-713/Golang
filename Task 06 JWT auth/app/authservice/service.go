package authservice

import (
	"fmt"
	"log"
	"net/http"
	"task6MuxGorm/models"
	"task6MuxGorm/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// auth service will provide // authorization middleware for user and admin //
type Service interface {
	VerifyJwtToken(tokenStr string) (*models.Claims, error)
	CreateJwtToken(user *models.User) (string, error)
	IsUser(parent http.HandlerFunc) http.HandlerFunc
	IsAdmin(parent http.HandlerFunc) http.HandlerFunc
}

type service struct {
	JWT_SECRET string
	DB         *gorm.DB
}

func New(JWT_SECRET string, DB *gorm.DB) Service {
	return &service{
		JWT_SECRET: JWT_SECRET,
		DB:         DB,
	}
}

// create auth service and give middlewares to respective one // this service requires secreat key

func (s *service) VerifyJwtToken(tokenStr string) (*models.Claims, error) {
	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid { // token is not valid so return status invalid
		return nil, err
	}

	return claims, nil
}

func (s *service) CreateJwtToken(user *models.User) (string, error) {

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &models.Claims{
		UserId: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{expirationTime},
		},
	}

	// create token from it

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// from this token get token string
	// ! key is of invalid type error returns if we don't give key input in byte array ha ha
	tokenString, err := token.SignedString([]byte(s.JWT_SECRET))

	return tokenString, err
}

func (s *service) IsAdmin(parent http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authorise user

		// get token from cookie
		token, err := r.Cookie("token")
		if err != nil {
			utils.JSONError(w, utils.ERROR_FINDING_COOKIE)
			return
		}

		// verify it
		claims, err := s.VerifyJwtToken(token.Value)
		if err != nil {
			utils.JSONError(w, &models.ApiError{Code: 400, Message: "Not valid token"})
			return
		}

		// check if he is admin
		if claims.Role != "admin" {
			utils.JSONError(w, &models.ApiError{Code: 400, Message: "Not admin"})
			return
		}

		// else continue

		parent.ServeHTTP(w, r)
		fmt.Println("Running after handler")
	})
}

func (s *service) IsUser(parent http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authorise user

		// get token from cookie
		token, err := r.Cookie("token")
		if err != nil {
			utils.JSONError(w, utils.ERROR_FINDING_COOKIE)
			return
		}

		log.Println("TOKEN AND ERROR ", token, err)

		// verify it
		claims, err := s.VerifyJwtToken(token.Value)
		if err != nil {
			log.Println(err)
			utils.JSONError(w, &models.ApiError{Code: 400, Message: "Not valid token"})
			return
		}

		// check if he is not user and admin then give error
		if claims.Role != "user" && claims.Role != "admin" {
			utils.JSONError(w, &models.ApiError{Code: 400, Message: "Not admin"})
			return
		}

		// else continue

		parent.ServeHTTP(w, r)
		fmt.Println("Running after handler")
	})
}

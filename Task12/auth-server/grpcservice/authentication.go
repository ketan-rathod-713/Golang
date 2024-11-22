package grpcservice

import (
	"auth/app/jwtauth"
	"auth/models"
	"context"
	"errors"
	"fmt"
	"log"
	authGrpc "proto/auth-server/v1"

	"github.com/golang-jwt/jwt/v5"
)

type AuthServer struct {
	Config     *models.Config
	JwtService jwtauth.Service
	authGrpc.AuthServer
}

func New(config *models.Config) *AuthServer {
	return &AuthServer{
		Config:     config,
		JwtService: jwtauth.New(),
	}
}

func (s *AuthServer) AuthoriseUser(ctx context.Context, req *authGrpc.AuthoriseRequest) (*authGrpc.AuthoriseResponse, error) {
	log.Println("GRPC AUTHORISE USER")

	if req.JwtToken == "" {
		return nil, errors.New("token is not present")
	}

	claims, err := s.JwtService.VerifyJwt(req.JwtToken, s.Config)

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			fmt.Println("JWT signature is invalid")
			return nil, errors.New("invalid jwt token")
		}

		fmt.Println("Error fetching jwt token")
		return nil, fmt.Errorf("error fetching JWT token: %s", err.Error())
	}

	// Verify from database that the given user exists
	// it must not be deleted or having another data
	// todo

	var response authGrpc.AuthoriseResponse = authGrpc.AuthoriseResponse{
		Email:    claims.Email,
		Role:     claims.Role,
		ObjectId: claims.ObjectId.Hex(),
	}

	fmt.Println("GRPC AUTH RESPONSE", response)
	return &response, nil
}

func (s *AuthServer) GetUserDetails(context.Context, *authGrpc.UserDetailsRequest) (*authGrpc.UserDetailsResponse, error) {
	return nil, nil
}

func (s *AuthServer) BookIssue(context.Context, *authGrpc.BookIssueRequest) (*authGrpc.BookIssueResponse, error) {
	fmt.Println("Book issue karo bhai")

	bookIssueResponse := authGrpc.BookIssueResponse{
		Issued: true,
		Error:  "no error",
	}

	return &bookIssueResponse, nil
}

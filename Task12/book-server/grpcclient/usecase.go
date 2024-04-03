package grpcclient

import (
	"context"
	"fmt"
	authGrpc "proto/auth-server/v1"
)

// write function for authorising user

type authClient struct {
	Client authGrpc.AuthClient
}

func NewAuthUse(client authGrpc.AuthClient) *authClient {
	return &authClient{
		Client: client,
	}
}

// each time i am creating and closing connection for now.
func (a *authClient) AuthoriseUser(token string) (*authGrpc.AuthoriseResponse, error) {
	// conn, client, err := NewAuthGrpcClient("localhost:8081")
	// defer conn.Close()

	// if err != nil {
	// 	fmt.Println("Error creating new auth grpc client", err)
	// 	return nil, err
	// }
	response, err := a.Client.AuthoriseUser(context.TODO(), &authGrpc.AuthoriseRequest{
		JwtToken: token,
	})

	if err != nil {
		fmt.Println("Error Occured In GRPC server", err)
		return nil, err
	}

	return response, nil
}

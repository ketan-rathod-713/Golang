package grpcclient

import (
	"fmt"
	authGrpcClient "proto/auth-server/v1"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthGrpcClient(host string) (*grpc.ClientConn, authGrpcClient.AuthClient, error) {
	var conn *grpc.ClientConn
	var err error

	for {
		// Attempt to establish connection with gRPC server
		fmt.Println("Attempting to connect grpc service")
		conn, err = grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		time.Sleep(1 * time.Second)
		if err != nil {
			fmt.Println("Failed to connect grpc service, attempting to connect in 5 seconds")
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}

	client := authGrpcClient.NewAuthClient(conn)
	if conn.GetState() == connectivity.Ready {
		fmt.Println("Auth grpc client connected")
	} else {
		fmt.Println("auth grpc client NOT connected")
	}

	return conn, client, nil
}

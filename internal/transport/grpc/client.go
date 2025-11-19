package grpc

import (
	"log"

	authv1 "github.com/shinoda4/sd-grpc-proto/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient(url string) (authv1.AuthServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := authv1.NewAuthServiceClient(conn)
	return client, conn
}

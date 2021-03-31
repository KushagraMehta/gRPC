package main

import (
	"context"
	"log"
	"time"

	pb "github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/protobuf/user"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := &pb.RegisterUserRequest{
		Fname:   "Steve",
		City:    "LA",
		Phone:   12345678,
		Height:  5.2,
		Married: true,
	}
	data, err := c.RegisterUser(ctx, user)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(data)
}

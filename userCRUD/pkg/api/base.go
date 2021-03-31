package api

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/protobuf/user"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	DB *pgxpool.Pool
	pb.UnimplementedUserServiceServer
}

func (server *Server) Initialize() {

	var err error

	DBURL := os.Getenv("DATABASE_URL")
	if server.DB, err = pgxpool.Connect(context.Background(), DBURL); err != nil {
		log.Println("Cannot connect to database")
		log.Fatal("This is the error:", err)
	} else {
		log.Println("We are connected to the database")
	}
	// log.Println(server.DB.Config())
}

func (server *Server) Run(port string) {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening to port %s\n", port)
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, server)
	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Println("Stopping the server")
	s.Stop()

	os.Exit(0)
}

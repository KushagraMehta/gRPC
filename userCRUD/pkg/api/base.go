package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

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
	var DBURL string
	if os.Getenv("LOCAL") == "1" {
		DBURL = fmt.Sprintf("postgres://%v:%v@%v:5432/crud", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), "localhostf")
	} else {
		DBURL = fmt.Sprintf("postgres://%v:%v@%v:5432/crud", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("HOST"))
	}
	for i := 1; i < 10; i++ {
		if server.DB, err = pgxpool.Connect(context.Background(), DBURL); err != nil {
			log.Println("Cannot connect to database")
			log.Printf("%v:Trying to connect to Database after 0.5Sec", i)
			time.Sleep(500 * time.Millisecond)
		} else {
			log.Println("We are connected to the database")
			break
		}
		if i == 4 {
			log.Fatal("This is the error:", err)
		}
	}
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

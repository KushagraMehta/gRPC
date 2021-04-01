package api

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/model"
	pb "github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/protobuf/user"
)

func (s *Server) RegisterUser(ctx context.Context, user *pb.RegisterUserRequest) (*pb.UserID, error) {
	log.Println("Registering User")
	uid, err := model.Register(ctx, s.DB, user)
	if err != nil {
		log.Println(err)
		return &pb.UserID{}, err
	}
	return &pb.UserID{ID: uid}, nil
}
func (s *Server) GetUserDetail(ctx context.Context, uid *pb.UserID) (*pb.UserDetailResponse, error) {

	log.Println("Fetching User Detail")
	user, err := model.FindUserByID(ctx, s.DB, uid.GetID())
	if err != nil {
		log.Println(err)
		return &pb.UserDetailResponse{}, err
	}
	return user, nil
}

func (s *Server) GetUsersList(ctx context.Context, uidList *pb.UsersListRequest) (*pb.UsersListResponse, error) {
	log.Println("Fetching Users list data ")
	uids := uidList.GetIDs()

	var data pb.UsersListResponse
	data.UserDetail = make([]*pb.UserDetailResponse, len(uids))

	for _, value := range uids {
		tmpUser, err := model.FindUserByID(ctx, s.DB, value)
		if err == nil {
			data.UserDetail = append(data.UserDetail, tmpUser)
		}
	}
	return &data, nil
}

func (s *Server) StreamUsersList(stream pb.UserService_StreamUsersListServer) error {
	log.Println("Streaming Users list data ")
	for {
		uid, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatalf("Error will streaming, %v", err)
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		data, err := model.FindUserByID(ctx, s.DB, uid.ID)
		if err != nil {
			return err
		}
		if err := stream.Send(data); err != nil {
			log.Fatalf("Error will send data,while streaming %v", err)
		}
	}

}

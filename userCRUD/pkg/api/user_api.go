package api

import (
	"context"
	"log"

	"github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/model"
	pb "github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/protobuf/user"
)

func (s *Server) RegisterUser(ctx context.Context, user *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	log.Println("Registering User")
	uid, err := model.Register(ctx, s.DB, user)
	if err != nil {
		log.Println(err)
		return &pb.RegisterUserResponse{}, err
	}
	return &pb.RegisterUserResponse{ID: uid}, nil
}
func (s *Server) GetUserDetail(ctx context.Context, uid *pb.UserDetailRequest) (*pb.UserDetailResponse, error) {

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

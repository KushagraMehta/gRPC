package model

import (
	"context"
	"errors"
	"regexp"
	"strconv"

	pb "github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/protobuf/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

// validate will check if User data has data or not
func validate(u *pb.RegisterUserRequest) error {
	var validNumber = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	if u.Fname == "" {
		return errors.New("please define a first name")
	}
	if u.City == "" {
		return errors.New("please define city")
	}
	if u.Phone == 0 && validNumber.MatchString(strconv.Itoa(int(u.Phone))) {
		return errors.New("please define a valid phone number")
	}
	if u.Height == 0 {
		return errors.New("please define height of the user")
	}
	return nil
}

// Register will store user data to database
func Register(ctx context.Context, db *pgxpool.Pool, u *pb.RegisterUserRequest) (uint32, error) {
	if err := validate(u); err != nil {
		return 0, err
	}
	var uid uint32
	if err := db.QueryRow(ctx, "INSERT INTO users(Fname,City,Phone,Height,Married) VALUES ($1,$2,$3,$4,$5) RETURNING id;", u.Fname, u.City, u.Phone, u.Height, u.Married).Scan(&uid); err != nil {
		return 0, err
	}
	return uid, nil
}

// FindUserByID will find a user with specific UID
func FindUserByID(ctx context.Context, db *pgxpool.Pool, uid uint32) (*pb.UserDetailResponse, error) {
	newUser := &pb.UserDetailResponse{}
	if err := db.QueryRow(ctx, "SELECT id,Fname,City,Phone,Height,Married FROM users WHERE ID=$1;", uid).Scan(&newUser.ID, &newUser.Fname, &newUser.City, &newUser.Phone, &newUser.Height, &newUser.Married); err != nil {
		return &pb.UserDetailResponse{}, err
	}
	return newUser, nil
}

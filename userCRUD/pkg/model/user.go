package model

import (
	"context"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"

	pb "github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/protobuf/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

// validate will check if User data has data or not
func validate(u *pb.RegisterUserRequest) error {
	var validNumber = regexp.MustCompile(`[0-9]{10}`)
	if u.Fname == "" {
		return errors.New("please define a first name")
	}
	if u.City == "" {
		return errors.New("please define city")
	}
	if u.Phone == 0 || !validNumber.MatchString(strconv.Itoa(int(u.Phone))) {
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
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_phone_key\" (SQLSTATE 23505)") {
			return 0, errors.New("phone number already register")
		} else {
			log.Println("Unknow error occure", err)
			return 0, errors.New("server internal error")
		}
	}
	return uid, nil
}

// FindUserByID will find a user with specific UID
func FindUserByID(ctx context.Context, db *pgxpool.Pool, uid uint32) (*pb.UserDetailResponse, error) {
	newUser := &pb.UserDetailResponse{}
	if err := db.QueryRow(ctx, "SELECT id,Fname,City,Phone,Height,Married FROM users WHERE ID=$1;", uid).Scan(&newUser.ID, &newUser.Fname, &newUser.City, &newUser.Phone, &newUser.Height, &newUser.Married); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return &pb.UserDetailResponse{}, errors.New("user does't exist")
		} else {
			log.Println("Unknow error occure", err)
			return &pb.UserDetailResponse{}, errors.New("server internal error")
		}
	}
	return newUser, nil
}

package model

import (
	"context"
	"math/rand"
	"strings"
	"testing"

	pb "github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/protobuf/user"
	"github.com/stretchr/testify/require"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func randomFloat() float32 {
	return float32(randomInt(2, 10)) * rand.Float32()
}
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func createRandomUser(t *testing.T) *pb.UserDetailResponse {
	testUser := pb.RegisterUserRequest{
		Fname:   randomString(5),
		City:    randomString(5),
		Phone:   randomInt(1000000000, 9999999999),
		Height:  randomFloat(),
		Married: randomInt(0, 1) != 0,
	}
	uid, err := Register(context.Background(), testDB.db, &testUser)

	require.NotZero(t, uid)
	require.NoError(t, err)

	return &pb.UserDetailResponse{
		ID:      uid,
		Fname:   testUser.Fname,
		City:    testUser.City,
		Phone:   testUser.Phone,
		Height:  testUser.Height,
		Married: testUser.Married,
	}
}

func TestRegister(t *testing.T) {
	var tests = []struct {
		input *pb.RegisterUserRequest
		err   string
	}{

		{&pb.RegisterUserRequest{
			Fname:   randomString(5),
			City:    randomString(5),
			Phone:   createRandomUser(t).Phone,
			Height:  randomFloat(),
			Married: randomInt(0, 1) != 0}, "phone number already register"},
		{&pb.RegisterUserRequest{
			Fname:   randomString(5),
			City:    randomString(5),
			Phone:   randomInt(1000000000, 9999999999),
			Height:  randomFloat(),
			Married: randomInt(0, 1) != 0}, ""},
		{&pb.RegisterUserRequest{
			Fname:   "",
			City:    randomString(5),
			Phone:   randomInt(1000000000, 9999999999),
			Height:  randomFloat(),
			Married: randomInt(0, 1) != 0}, "please define a first name"},
		{&pb.RegisterUserRequest{
			Fname:   randomString(5),
			City:    "",
			Phone:   randomInt(1000000000, 9999999999),
			Height:  randomFloat(),
			Married: randomInt(0, 1) != 0}, "please define city"},
		{&pb.RegisterUserRequest{
			Fname:   randomString(5),
			City:    randomString(5),
			Phone:   0,
			Height:  randomFloat(),
			Married: randomInt(0, 1) != 0}, "please define a valid phone number"},
		{&pb.RegisterUserRequest{
			Fname:   randomString(5),
			City:    randomString(5),
			Phone:   123456789,
			Height:  randomFloat(),
			Married: randomInt(0, 1) != 0}, "please define a valid phone number"},
	}
	for _, test := range tests {
		uid, err := Register(context.Background(), testDB.db, test.input)
		if test.err != "" {
			require.EqualError(t, err, test.err)
		} else {
			require.NotZero(t, uid)
			require.Nil(t, err)
		}
		cleanTable(t)
	}
}

func TestFindUserByID(t *testing.T) {
	t.Run("When User exist", func(t *testing.T) {

		userData := createRandomUser(t)

		output, err := FindUserByID(context.Background(), testDB.db, userData.ID)
		require.NoError(t, err)
		require.Equal(t, userData.Fname, output.Fname)
		require.Equal(t, userData.City, output.City)
		require.Equal(t, userData.Height, output.Height)
		require.Equal(t, userData.Married, output.Married)
		require.Equal(t, userData.Phone, output.Phone)
	})
	t.Run("When User does't exist", func(t *testing.T) {

		_, err := FindUserByID(context.Background(), testDB.db, 0)
		require.EqualError(t, err, "user does't exist")
	})
	cleanTable(t)
}

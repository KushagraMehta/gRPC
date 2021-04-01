package model

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

type DB struct {
	db *pgxpool.Pool
}

var testDB DB

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("sad .env file Not found")
	}
}
func cleanTable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := testDB.db.Exec(ctx, "delete from users;")
	require.NoError(t, err)
}

func TestMain(m *testing.M) {
	var err error
	DBURL := os.Getenv("DATABASE_URL")
	if testDB.db, err = pgxpool.Connect(context.Background(), DBURL); err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	os.Exit(m.Run())
}

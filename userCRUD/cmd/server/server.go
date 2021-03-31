package main

import (
	"log"
	"os"

	api "github.com/KushagraMehta/gRPC-Tutorial/userCRUD/pkg/api"
	"github.com/joho/godotenv"
)

const PORT = ":50051"

var server = api.Server{}

func init() {
	if os.Getenv("LOCAL") == "1" {
		log.Print("Running Locally")
		if err := godotenv.Load(); err != nil {
			log.Print("sad .env file not found")
		} else {
			log.Print("We are getting the env values")
		}
	} else {

		log.Print("Running on server")
	}
}
func main() {
	server.Initialize()
	server.Run(":" + os.Getenv("PORT"))
	defer server.DB.Close()
}

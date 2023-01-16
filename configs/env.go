package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURL() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}

	return os.Getenv("MONGOURI")
}

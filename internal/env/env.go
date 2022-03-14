package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func SetupEnv() {
	var env string

	if os.Getenv("DOCKER") == "true" {
		env = ".env.dev"
	} else {
		env = "internal/env/.env.dev"
	}
	err := godotenv.Load(env)

	if err != nil {
		log.Fatalf("Variable environment not found")
	}
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".jwt_env")
	if err != nil {
		log.Fatal("Error loading .jwt_env file")
	}
}

var (
	SecretKey = os.Getenv("JWT_SECRET")
)

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".jwt_service")
	if err != nil {
		log.Fatal("Error loading .jwt_service file")
	}
}

var (
	JwtSecret = os.Getenv("JWT_SECRET")
)
